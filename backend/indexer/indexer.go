package indexer

import (
	"context"
	"time"

	"github.com/D-pixel-crime/Freelance_Escrow/backend/utils"
	"github.com/charmbracelet/log"
)

func StartIndexer() {
	log.Infof("Indexer started. Waiting for blocks...")

	// Simple polling loop to print latest block number
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		if utils.Web3Client == nil {
			log.Errorf("Web3Client is nil, indexer cannot proceed.")
			return
		}
		blockNumber, err := utils.Web3Client.BlockNumber(context.Background())
		if err != nil {
			log.Errorf("Failed to retrieve latest block: %v", err)
			continue
		}
		log.Infof("Indexer: Successfully connected. Latest block number: %d", blockNumber)
	}
}
