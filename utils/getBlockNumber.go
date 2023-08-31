package utils

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetNowBlockNumberAndBlockTime(client *ethclient.Client) (uint64, uint64) {
	// Get the latest known block

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		return 0, 0
	}
	blockNumber := block.NumberU64()
	blockTime := block.Time()
	return blockNumber, blockTime

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
