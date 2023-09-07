package utils

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TransferETH(client *ethclient.Client, privateKey *ecdsa.PrivateKey,
	to string, amount string) {

	account := GetAccount(privateKey)
	//获取nonce
	nonce := GetAccountNonce(client, account)

	toAddress := common.HexToAddress(to)

	value := StringToBig(amount)

	gasLimit := uint64(21000)

	gasPrice, _ := client.SuggestGasPrice(context.Background())

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	chainID, _ := client.NetworkID(context.Background())

	signedTx, _ := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)

	err := client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx Hash: %s", signedTx.Hash().Hex())
}
