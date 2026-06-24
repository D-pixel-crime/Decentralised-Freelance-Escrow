package indexer

import (
	"context"
	"os"
	"time"

	"github.com/D-pixel-crime/Freelance_Escrow/backend/contracts/escrow"
	"github.com/D-pixel-crime/Freelance_Escrow/backend/utils"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func StartIndexer(jobsColl *mongo.Collection) {
	log.Infof("Indexer started. Waiting for blocks & events...")

	if utils.Web3Client == nil {
		log.Errorf("Web3Client is nil, indexer cannot proceed.")
		return
	}

	contractAddressHex := os.Getenv("ESCROW_CONTRACT_ADDRESS")
	if contractAddressHex == "" {
		log.Errorf("ESCROW_CONTRACT_ADDRESS not set, indexer exiting.")
		return
	}

	contractAddress := common.HexToAddress(contractAddressHex)
	escrowInstance, err := escrow.NewFreelanceEscrow(contractAddress, utils.Web3Client)
	if err != nil {
		log.Errorf("Failed to instantiate FreelanceEscrow contract: %v", err)
		return
	}

	sink := make(chan *escrow.FreelanceEscrowFreelanceEscrowClientStakeCompleted)
	sub, err := escrowInstance.WatchFreelanceEscrowClientStakeCompleted(&bind.WatchOpts{Context: context.Background()}, sink, nil, nil)
	if err != nil {
		log.Errorf("Failed to subscribe to ClientStakeCompleted event: %v", err)
		return
	}
	defer sub.Unsubscribe()

	log.Infof("Successfully subscribed to FreelanceEscrow events at %s", contractAddressHex)

	for {
		select {
		case err := <-sub.Err():
			log.Errorf("Event subscription error: %v", err)
			return
		case event := <-sink:
			log.Infof("Detected ClientStakeCompleted! Amount: %s, Timestamp: %s", event.Amount.String(), event.Timestamp.String())
			
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			
			filter := bson.M{"contractAddress": contractAddressHex}
			update := bson.M{
				"$set": bson.M{
					"status": "CLIENT_STAKED",
					"clientStakedAmount": event.Amount.String(),
				},
			}
			
			res, err := jobsColl.UpdateOne(ctx, filter, update)
			cancel()
			
			if err != nil {
				log.Errorf("Error updating MongoDB for Job at %s: %v", contractAddressHex, err)
			} else if res.MatchedCount == 0 {
				log.Warnf("No MongoDB Job found matching contract address %s", contractAddressHex)
			} else {
				log.Infof("Successfully updated MongoDB Job %s to CLIENT_STAKED", contractAddressHex)
			}
		}
	}
}
