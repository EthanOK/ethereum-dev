package utils

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetLatestBlockNumberAndBlockTime(client *ethclient.Client) (uint64, uint64) {
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
func GetLatestBlockNumber(client *ethclient.Client) *big.Int {
	// Get the latest known block

	block, _ := client.BlockByNumber(context.Background(), nil)

	blockNumber := block.Number()

	return blockNumber

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

// GetAccountNonce
func GetAccountNonce(client *ethclient.Client, account string) uint64 {
	address := common.HexToAddress(account)
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		log.Fatal(err)
	}
	return nonce
}
