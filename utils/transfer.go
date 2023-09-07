package utils

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"gocode.ethan/ethereum-dev/utils"
)

func TransferETH(client *ethclient.Client, privateKey *ecdsa.PrivateKey, toAddress string, value *big.Int) {
	//获取nonce
	nonce := utils.GetAccountNonce(client, account)
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
}
