package posthandlers

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/D-pixel-crime/Freelance_Escrow/backend/contracts"
	"github.com/D-pixel-crime/Freelance_Escrow/backend/models"
	"github.com/D-pixel-crime/Freelance_Escrow/backend/utils"
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
	ClientEthAccount     string        `json:"clientEthAccount" binding:"required"`
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

func deployContract(freelancerEthAccount, clientEthAccount string, mongoJobId bson.ObjectID) (string, error) {
	client, err := ethclient.Dial(os.Getenv("SEPOLIA_RPC_URL"))
	if err != nil {
		return "", fmt.Errorf("%s", "Error Connecting with Ethereum-Endpoint!"+err.Error())
	}

	contractOwner, err := crypto.HexToECDSA(os.Getenv("DEPLOYER_ACCOUNT"))
	if err != nil {
		return "", err
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return "", err
	}

	authenticatedTransactor, err := bind.NewKeyedTransactorWithChainID(contractOwner, chainId)
	if err != nil {
		return "", err
	}

	jobId := new(big.Int).SetBytes(mongoJobId[:])
	clientAddr := common.HexToAddress(clientEthAccount)
	freelancerAddr := common.HexToAddress(freelancerEthAccount)
	arbitratorAddr := common.HexToAddress(os.Getenv("ARBITRATOR_ACCOUNT"))
	confirmationPeriod := big.NewInt(int64(time.Hour * 24 * 2))

	contractAddr, _, _, err := contracts.DeployFreelanceEscrow(authenticatedTransactor, client, jobId, clientAddr, freelancerAddr, arbitratorAddr, confirmationPeriod)
	if err != nil {
		return "", err
	}

	return contractAddr.Hex(), nil
}

func updateJob(freelancerEthAccount, contractAddr string, jobId bson.ObjectID) error {
	var res models.Job
	coll := utils.DBClient.Database(os.Getenv("DATABASE_NAME")).Collection("jobs")
	filter := bson.M{"_id": jobId}

	freelancerId, err := checkFreelancer(freelancerEthAccount)
	if err != nil {
		return err
	}

	update := bson.M{"$set": bson.M{"freelancerId": freelancerId, "status": models.AGREED, "contractAddress": contractAddr}}
	err = coll.FindOneAndUpdate(context.TODO(), filter, update).Decode(&res)
	if err != nil {
		return err
	}

	return nil
}

func AllocateJob(c *gin.Context) {
	var reqBody jobAllocationRequest
	var err error

	if err = c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect Request Format!"})
		return
	}

	contractAddr, err := deployContract(reqBody.FreelancerEthAccount, reqBody.ClientEthAccount, reqBody.JobId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error!" + err.Error()})
		return
	}

	err = updateJob(reqBody.FreelancerEthAccount, contractAddr, reqBody.JobId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "User Not Found!"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error!"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Job Allocated!"})
}
