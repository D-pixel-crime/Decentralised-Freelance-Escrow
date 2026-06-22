package utils

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/ethclient"
)

var Web3Client *ethclient.Client

func ConnectToWeb3() (*ethclient.Client, error) {
	rpcUrl := os.Getenv("WEB3_RPC_URL")
	if rpcUrl == "" {
		return nil, fmt.Errorf("WEB3_RPC_URL environment variable is missing")
	}

	// Dial to Web3 Provider
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, fmt.Errorf("Error Connecting to Web3! Error:%s", err)
	}

	Web3Client = client
	log.Infof("Web3 Connection Successful: %s", rpcUrl)

	return Web3Client, nil
}
