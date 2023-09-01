package main

import (
	"fmt"

	"gocode.ethan/ethereum-dev/utils"
)

func main() {

	/* 	client, err := getclient.GetEthClient()
	   	if err != nil {
	   		log.Fatal(err)
	   		return
	   	}
	   	account := "0x95B3Aad7f20E78a0e4DcEB9c23beca4e55ebdDF6"
	   	// Get the balance of an account
	   	balance := utils.GetETHBalance(client, account)
	   	// fmt.Println("Account balance:", balance, "WEI")
	   	fmt.Println("Account balance:", utils.WeiToEther(balance), "ETH")
	   	// Get the latest known block
	   	blockNumber, blockTime := utils.GetNowBlockNumberAndBlockTime(client)
	   	fmt.Println("Latest Block Number:", blockNumber)
	   	fmt.Println("Latest Block Time:", blockTime)
	   	// Get the balance of USDT
	   	contractUSDT := "0xdAC17F958D2ee523a2206206994597C13D831ec7"
	   	metadata := utils.GetERC20TokenMetadata(client, contractUSDT)
	   	symbol := metadata.Symbol
	   	decimals := metadata.Decimals
	   	usdtBalance := utils.GetERC20Balance(client, contractUSDT, account)
	   	fmt.Println("Account balance:", utils.BigToDecimals(usdtBalance,decimals), symbol)
	*/

	privateKey := utils.GeneratePrivateKeyHasPrefix()
	address := utils.PrivateKeyToAddress(privateKey)
	isAddress := utils.CheckIsAddress(address)
	fmt.Println(address, "is Address: ", isAddress)

}
