package gethandlers

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

// private functions

func getAddressBalance(addr string, rpcUrl string) (*big.Float, error) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, err
	}

	account := common.HexToAddress(addr)

	balanceWei, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return nil, err
	}

	fBalance, _ := new(big.Float).SetString(balanceWei.String())
	ethValue := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))

	return ethValue, nil
}

// public functions

func GetWalletBalance(c *gin.Context) {
	address := c.Param("address")
	balance, err := getAddressBalance(address, os.Getenv("SEPOLIA_RPC_URL"))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch balance"})
		return
	}

	c.JSON(200, gin.H{
		"address": address,
		"balance": fmt.Sprintf("%f ETH", balance),
	})
}
