package utils

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gocode.ethan/ethereum-dev/token/erc20"
)

func GetETHBalance(client *ethclient.Client, address string) *big.Int {

	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return balance
}

func GetERC20Balance(client *ethclient.Client, contract string, account string) *big.Int {

	contractAddress := common.HexToAddress(contract)
	accountAddress := common.HexToAddress(account)
	// tokenInstance, err := token.NewToken(tokenAddress, client)
	tokenInstance, err := erc20.NewERC20(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// &bind.CallOpts{} 等价于 signer
	balance, err := tokenInstance.BalanceOf(nil, accountAddress)
	if err != nil {
		log.Fatal(err)
	}
	return balance
}

type TokenMetadata struct {
	Name     string
	Symbol   string
	Decimals uint8
}

func GetERC20TokenMetadata(client *ethclient.Client, contract string) TokenMetadata {
	contractAddress := common.HexToAddress(contract)
	// tokenInstance, err := token.NewToken(tokenAddress, client)
	tokenInstance, err := erc20.NewERC20(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	name, _ := tokenInstance.Name(nil)
	symbol, _ := tokenInstance.Symbol(nil)
	decimals, _ := tokenInstance.Decimals(nil)
	metadata := TokenMetadata{name, symbol, decimals}
	return metadata
}
