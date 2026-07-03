package posthandlers

import (
	"context"
	"fmt"
	"math/big"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/D-pixel-crime/Freelance_Escrow/backend/contracts"
	"github.com/D-pixel-crime/Freelance_Escrow/backend/models"
	"github.com/D-pixel-crime/Freelance_Escrow/backend/utils"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type jobAllocationRequest struct {
	FreelancerEthAccount string        `json:"freelancerEthAccount" binding:"required"`
	JobId                bson.ObjectID `json:"jobId" binding:"required"`
	ChainId              int           `json:"chainId" binding:"required"`
}

func checkFreelancer(freelancerEthAccount string) (bson.ObjectID, error) {
	var res models.Freelancer
	coll := utils.DBClient.Database(os.Getenv("DATABASE_NAME")).Collection("freelancer")
	filter := bson.M{"ethAccount": freelancerEthAccount}

	err := coll.FindOne(context.TODO(), filter).Decode(&res)
	if err != nil {
		return bson.ObjectID{}, err
	}

	return res.ID, nil
}

func deployContract(freelancerEthAccount, clientEthAccount string, mongoJobId bson.ObjectID) (string, string, error) {
	rpcURL := os.Getenv("WEB3_RPC_URL")
	if rpcURL == "" {
		rpcURL = "http://127.0.0.1:8545"
	}

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return "", "", fmt.Errorf("Failed to connect to Ethereum RPC at %s: %v", rpcURL, err)
	}

	privKeyStr := os.Getenv("PRIVATE_KEY")
	if privKeyStr == "" {
		privKeyStr = os.Getenv("DEPLOYER_ACCOUNT")
	}
	if privKeyStr == "" {
		// Hardcoded fallback: Anvil Account 0 private key for local development
		privKeyStr = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
		log.Warnf("PRIVATE_KEY env var is empty — using hardcoded Anvil Account 0 fallback")
	}
	privKeyStr = strings.TrimPrefix(privKeyStr, "0x")

	contractOwner, err := crypto.HexToECDSA(privKeyStr)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse private key: %v", err)
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return "", "", err
	}

	authenticatedTransactor, err := bind.NewKeyedTransactorWithChainID(contractOwner, chainId)
	if err != nil {
		return "", "", fmt.Errorf("Failed to create transactor: %v", err)
	}

	jobId := new(big.Int).SetBytes(mongoJobId[:])
	clientAddr := common.HexToAddress(clientEthAccount)
	freelancerAddr := common.HexToAddress(freelancerEthAccount)
	
	var arbitrators []models.Arbitrator
	arbColl := utils.DBClient.Database(os.Getenv("DATABASE_NAME")).Collection("arbitrator")
	cursor, err := arbColl.Find(context.TODO(), bson.M{})
	
	arbitratorAccountStr := os.Getenv("ARBITRATOR_ACCOUNT")
	if err == nil {
		if err = cursor.All(context.TODO(), &arbitrators); err == nil && len(arbitrators) > 0 {
			randSource := rand.New(rand.NewSource(time.Now().UnixNano()))
			selected := arbitrators[randSource.Intn(len(arbitrators))]
			arbitratorAccountStr = selected.EthAccount
		}
	}
	arbitratorAddr := common.HexToAddress(arbitratorAccountStr)
	
	confirmationPeriod := big.NewInt(int64(time.Hour * 24 * 2))

	contractAddr, _, _, err := contracts.DeployFreelanceEscrow(authenticatedTransactor, client, jobId, clientAddr, freelancerAddr, arbitratorAddr, confirmationPeriod)
	if err != nil {
		return "", "", err
	}

	return contractAddr.Hex(), arbitratorAccountStr, nil
}

func updateJob(freelancerEthAccount, contractAddr string, jobId bson.ObjectID, arbitratorAccountStr string) error {
	var res models.Job
	coll := utils.DBClient.Database(os.Getenv("DATABASE_NAME")).Collection("jobs")
	filter := bson.M{"_id": jobId}

	freelancerId, err := checkFreelancer(freelancerEthAccount)
	if err != nil {
		return err
	}

	update := bson.M{"$set": bson.M{"freelancerId": freelancerId, "status": models.AGREED, "contractAddress": contractAddr, "arbitratorEth": arbitratorAccountStr}}
	err = coll.FindOneAndUpdate(context.TODO(), filter, update).Decode(&res)
	if err != nil {
		return err
	}

	return nil
}

func getClientEthAccount(jobId bson.ObjectID) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var job models.Job
	jobsColl := utils.DBClient.Database(os.Getenv("DATABASE_NAME")).Collection("jobs")
	if err := jobsColl.FindOne(ctx, bson.M{"_id": jobId}).Decode(&job); err != nil {
		return "", err
	}

	var client models.Client
	clientColl := utils.DBClient.Database(os.Getenv("DATABASE_NAME")).Collection("client")
	if err := clientColl.FindOne(ctx, bson.M{"_id": job.ClientID}).Decode(&client); err != nil {
		return "", err
	}

	return client.EthAccount, nil
}

func AllocateJob(c *gin.Context) {
	var reqBody jobAllocationRequest
	var err error

	if err = c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect Request Format!"})
		return
	}

	clientEthAccount, err := getClientEthAccount(reqBody.JobId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Job or Client Not Found!"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error! " + err.Error()})
		}
		return
	}

	contractAddr, arbitratorAccountStr, err := deployContract(reqBody.FreelancerEthAccount, clientEthAccount, reqBody.JobId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error!" + err.Error()})
		return
	}

	err = updateJob(reqBody.FreelancerEthAccount, contractAddr, reqBody.JobId, arbitratorAccountStr)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "User Not Found!"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error!"})
		}
		return
	}

	log.Infof("Job %s allocated → Contract deployed at %s and saved to MongoDB", reqBody.JobId.Hex(), contractAddr)
	c.JSON(http.StatusOK, gin.H{"message": "Job Allocated!", "contractAddress": contractAddr})
}
