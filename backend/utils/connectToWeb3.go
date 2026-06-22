package utils

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/ethclient"
)

var Web3Client *ethclient.Client

func ConnectToWeb3() (*ethclient.Client, error) {
	// Dial to local Anvil node
	client, err := ethclient.Dial("ws://127.0.0.1:8545")
	if err != nil {
		return nil, fmt.Errorf("Error Connecting to Web3! Error:%s", err)
	}

	Web3Client = client
	log.Infof("Web3 Connection Successful: ws://127.0.0.1:8545")

	return Web3Client, nil
}
