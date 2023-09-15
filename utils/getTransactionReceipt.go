package utils

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetTransactionReceipt(client *ethclient.Client, txHex string) *types.Receipt {
	// 获取交易的回执
	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(txHex))
	if err != nil {
		log.Fatal(err)
	}
	return receipt
}
