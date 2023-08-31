package utils

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetNowBlockNumberTime(client *ethclient.Client) *types.Block {
	// Get the latest known block

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return block

}

// GetLatestBlock
func GetLatestBlock(client *ethclient.Client) *types.Block {
	// Get the latest known block

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return block
}
