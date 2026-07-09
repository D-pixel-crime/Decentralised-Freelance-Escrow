package gethandlers

import (
	"context"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/D-pixel-crime/Freelance_Escrow/backend/contracts/escrow"
	"github.com/D-pixel-crime/Freelance_Escrow/backend/models"
	"github.com/D-pixel-crime/Freelance_Escrow/backend/utils"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// mapEscrowStateToStatus maps the on-chain uint8 enum to our application's JobStatus string.
// This mirrors the Solidity enum ordering in the FreelanceEscrow contract.
func mapEscrowStateToStatus(state uint8) models.JobStatus {
	switch state {
	case 0:
		return models.AGREED
	case 1:
		return models.CLIENT_STAKED
	case 2:
		return models.FREELANCER_STAKED
	case 3:
		return models.ALL_STAKED_AND_PENDING
	case 4:
		return models.PENDING_CLIENT_CONFIRMATION
	case 5:
		return models.JOB_COMPLETED
	case 6:
		return models.CANCEL_REQUESTED
	case 7:
		return models.DEAL_BROKEN
	case 8:
		return models.RANDOM_DISPUTED
	case 9:
		return models.PAYMENT_DISPUTED
	default:
		return "UNKNOWN_STATE"
	}
}

// isTerminalState returns true for states where the escrow is finalized
// and no further on-chain transitions are expected, or the job has no contract.
func isTerminalState(status models.JobStatus) bool {
	return status == models.JOB_COMPLETED || status == models.DEAL_BROKEN || status == models.UNALLOCATED
}

// jitSyncJobs performs Just-In-Time synchronization of on-chain escrow state
// for all active (non-terminal) jobs that have a deployed contract.
// It connects to the local Anvil RPC, reads GetEscrowState() for each job,
// and updates MongoDB + the in-memory slice if a drift is detected.
func jitSyncJobs(jobs []models.Job, dbName string) []models.Job {
	const rpcURL = "http://127.0.0.1:8545"

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Errorf("[JIT-Sync] Failed to connect to RPC at %s: %v", rpcURL, err)
		return jobs // Return unmodified — graceful degradation
	}
	defer client.Close()

	jobsColl := utils.DBClient.Database(dbName).Collection("jobs")

	for i := range jobs {
		job := &jobs[i]

		// Skip jobs without a deployed contract
		if job.ContractAddress == "" {
			continue
		}

		// Skip terminal states — no further transitions possible
		if isTerminalState(job.Status) {
			continue
		}

		// Bind to the contract and read on-chain state
		contractAddr := common.HexToAddress(job.ContractAddress)
		escrowInstance, err := escrow.NewFreelanceEscrow(contractAddr, client)
		if err != nil {
			log.Errorf("[JIT-Sync] Failed to bind contract %s: %v", job.ContractAddress, err)
			continue
		}

		callCtx, callCancel := context.WithTimeout(context.Background(), 3*time.Second)
		stateUint, err := escrowInstance.GetEscrowState(&bind.CallOpts{Context: callCtx})
		callCancel()

		if err != nil {
			// Detect stale contracts from a restarted Anvil session
			if strings.Contains(err.Error(), "no contract code at given address") {
				log.Warnf("[JIT-Sync] Stale contract %s — no code on-chain (Anvil likely restarted). Skipping.", job.ContractAddress)
			} else {
				log.Errorf("[JIT-Sync] Failed to read EscrowState for %s: %v", job.ContractAddress, err)
			}
			continue
		}

		onChainStatus := mapEscrowStateToStatus(stateUint)

		// If MongoDB is already in sync, skip the write
		if job.Status == onChainStatus {
			continue
		}

		// Drift detected — update MongoDB to match the chain
		log.Infof("[JIT-Sync] Drift detected for contract %s: DB=%s → Chain=%s. Updating...",
			job.ContractAddress, job.Status, onChainStatus)

		updateCtx, updateCancel := context.WithTimeout(context.Background(), 3*time.Second)
		filter := bson.M{"_id": job.ID}
		update := bson.M{"$set": bson.M{"status": onChainStatus}}

		_, err = jobsColl.UpdateOne(updateCtx, filter, update)
		updateCancel()

		if err != nil {
			log.Errorf("[JIT-Sync] Failed to update MongoDB for job %s: %v", job.ID.Hex(), err)
			continue
		}

		// Update the in-memory job so the API response reflects the corrected state
		job.Status = onChainStatus
		log.Infof("[JIT-Sync] Successfully synced job %s to %s", job.ID.Hex(), onChainStatus)
	}

	return jobs
}

func GetMyJobs(c *gin.Context) {
	roleStr := c.GetString("role")
	ethAccountStr := c.GetString("ethAccount")

	if roleStr == "" || ethAccountStr == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid context variables"})
		return
	}

	var userID bson.ObjectID
	dbName := os.Getenv("DATABASE_NAME")

	if roleStr == "client" || roleStr == "freelancer" {
		var user struct{ ID bson.ObjectID `bson:"_id"` }
		if err := utils.DBClient.Database(dbName).Collection(roleStr).FindOne(context.TODO(), bson.M{"ethAccount": ethAccountStr}).Decode(&user); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		userID = user.ID
	} else if roleStr != "arbitrator" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	coll := utils.DBClient.Database(dbName).Collection("jobs")
	var filter bson.M

	switch roleStr {
	case "client":
		filter = bson.M{"clientId": userID}
	case "arbitrator":
		filter = bson.M{"arbitratorEth": ethAccountStr}
	default:
		filter = bson.M{
			"$or": []bson.M{
				{"freelancerId": userID},
				{"applicants": ethAccountStr},
			},
		}
	}

	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query jobs"})
		return
	}
	defer cursor.Close(ctx)

	var jobs []models.Job
	if err := cursor.All(ctx, &jobs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode jobs"})
		return
	}

	if jobs == nil {
		jobs = []models.Job{}
	}

	jobs = jitSyncJobs(jobs, dbName)

	c.JSON(http.StatusOK, jobs)
}
