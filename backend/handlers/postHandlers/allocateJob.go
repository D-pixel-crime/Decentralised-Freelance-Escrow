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

	"github.com/D-pixel-crime/Freelance_Escrow/backend/contracts/factory"
	"github.com/D-pixel-crime/Freelance_Escrow/backend/models"
	"github.com/D-pixel-crime/Freelance_Escrow/backend/utils"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type jobAllocationRequest struct {
	FreelancerEthAccount string        `json:"freelancerEthAccount" binding:"required"`
	JobId                bson.ObjectID `json:"jobId" binding:"required"`
	ChainId              int           `json:"chainId" binding:"required"`
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

	authenticatedTransactor.GasLimit = uint64(3000000)

	factoryAddr := os.Getenv("FACTORY_ADDRESS")
	if factoryAddr == "" {
		return "", "", fmt.Errorf("FACTORY_ADDRESS env var is missing")
	}

	factoryInstance, err := factory.NewFactory(common.HexToAddress(factoryAddr), client)
	if err != nil {
		return "", "", err
	}

	tx, err := factoryInstance.CreateEscrow(authenticatedTransactor, jobId, clientAddr, freelancerAddr, arbitratorAddr, confirmationPeriod)
	if err != nil {
		return "", "", err
	}

	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return "", "", err
	}

	var childContractAddr string
	for _, vLog := range receipt.Logs {
		event, err := factoryInstance.ParseEscrowCreated(*vLog)
		if err == nil && event != nil {
			childContractAddr = event.EscrowAddress.Hex()
			break
		}
	}

	if childContractAddr == "" {
		return "", "", fmt.Errorf("EscrowCreated event not found in logs")
	}

	return childContractAddr, arbitratorAccountStr, nil
}

func AllocateJob(c *gin.Context) {
	var reqBody jobAllocationRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect Request Format!"})
		return
	}

	dbName := os.Getenv("DATABASE_NAME")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var job models.Job
	if err := utils.DBClient.Database(dbName).Collection("jobs").FindOne(ctx, bson.M{"_id": reqBody.JobId}).Decode(&job); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job Not Found!"})
		return
	}

	var client models.Client
	if err := utils.DBClient.Database(dbName).Collection("client").FindOne(ctx, bson.M{"_id": job.ClientID}).Decode(&client); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Client Not Found!"})
		return
	}

	var freelancer models.Freelancer
	if err := utils.DBClient.Database(dbName).Collection("freelancer").FindOne(ctx, bson.M{"ethAccount": reqBody.FreelancerEthAccount}).Decode(&freelancer); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Freelancer Not Found!"})
		return
	}

	contractAddr, arbitratorAccountStr, err := deployContract(reqBody.FreelancerEthAccount, client.EthAccount, reqBody.JobId)
	if err != nil {
		log.Errorf("Failed to deploy contract: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error! " + err.Error()})
		return
	}

	update := bson.M{
		"$set": bson.M{
			"freelancerId":    freelancer.ID,
			"status":          models.AGREED,
			"contractAddress": contractAddr,
			"arbitratorEth":   arbitratorAccountStr,
		},
	}
	if _, err := utils.DBClient.Database(dbName).Collection("jobs").UpdateOne(ctx, bson.M{"_id": reqBody.JobId}, update); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error!"})
		return
	}

	log.Infof("Job %s allocated → Contract deployed at %s and saved to MongoDB", reqBody.JobId.Hex(), contractAddr)
	c.JSON(http.StatusOK, gin.H{"message": "Job Allocated!", "contractAddress": contractAddr})
}
