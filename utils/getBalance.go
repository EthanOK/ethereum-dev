package utils

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gocode.ethan/ethereum-dev/token/erc20"
	"gocode.ethan/ethereum-dev/token/erc721"
)

// Get ETH Balance
func GetETHBalance(client *ethclient.Client, address string) *big.Int {

	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return balance
}

// return (balance, decimals)
func GetERC20Balance(client *ethclient.Client, contract string, account string) (*big.Int, uint8) {

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
	decimals, _ := tokenInstance.Decimals(nil)
	return balance, decimals
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

func GetERC20Decimals(client *ethclient.Client, contract string) uint8 {
	contractAddress := common.HexToAddress(contract)
	// tokenInstance, err := token.NewToken(tokenAddress, client)
	tokenInstance, err := erc20.NewERC20(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	decimals, _ := tokenInstance.Decimals(nil)
	return decimals
}

// return (balance, symbol)
func GetERC721Balance(client *ethclient.Client, contract string, account string) (*big.Int, string) {

	contractAddress := common.HexToAddress(contract)
	accountAddress := common.HexToAddress(account)
	// tokenInstance, err := token.NewToken(tokenAddress, client)
	tokenInstance, err := erc721.NewERC721(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// &bind.CallOpts{} 等价于 signer
	balance, err := tokenInstance.BalanceOf(nil, accountAddress)
	if err != nil {
		log.Fatal(err)
	}
	symbol, _ := tokenInstance.Symbol(nil)
	return balance, symbol
}
