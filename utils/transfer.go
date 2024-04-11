package utils

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"gocode.ethan/ethereum-dev/contracts/erc20"
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

	chainID, _ := client.ChainID(context.Background())

	signedTx, _ := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)

	err := client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx Hash: %s", signedTx.Hash().Hex())
}

func TransferERC20(client *ethclient.Client, privateKey *ecdsa.PrivateKey,
	contract string, to string, amount string) {

	contractAddress := common.HexToAddress(contract)
	toAddress := common.HexToAddress(to)
	amount_ := StringToBig(amount)
	chainId, _ := client.ChainID(context.Background())
	user, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	tokenInstance, err := erc20.NewERC20(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := tokenInstance.Transfer(user, toAddress, amount_)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
}
