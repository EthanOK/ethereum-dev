package main

import (
	"fmt"
	"log"

	"gocode.ethan/ethereum-dev/getclient"
	"gocode.ethan/ethereum-dev/utils"
)

func main() {

	client, err := getclient.GetEthClient()
	if err != nil {
		log.Fatal(err)
		return
	}
	address := "0x95B3Aad7f20E78a0e4DcEB9c23beca4e55ebdDF6"
	balance := utils.GetETHBalance(client, address)
	fmt.Println("Account balance:", balance, "WEI")
	// wei to eth
	fmt.Println("Account balance:", utils.WeiToEther(balance), "ETH")
	// Get the latest known block
	blockNumber, blockTime := utils.GetNowBlockNumberAndBlockTime(client)
	fmt.Println("Latest Block Number:", blockNumber)
	fmt.Println("Latest Block Time:", blockTime)

}
