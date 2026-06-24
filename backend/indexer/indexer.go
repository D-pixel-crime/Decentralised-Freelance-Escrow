package indexer

import (
	"context"
	"os"

	"github.com/D-pixel-crime/Freelance_Escrow/backend/contracts/escrow"
	"github.com/D-pixel-crime/Freelance_Escrow/backend/utils"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func StartIndexer() {
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
			log.Infof("Attempting to update MongoDB for Job/Escrow ID (Pending structured query...)")
		}
	}
}
