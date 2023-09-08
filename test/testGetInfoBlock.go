package test

import (
	"fmt"

	"gocode.ethan/ethereum-dev/getclient"
	"gocode.ethan/ethereum-dev/utils"
)

func TestGetInfoBlock() {
	// Get Client
	client, _ := getclient.GetEthClient()

	// Get CurrentBlockNumber CurrentBlockTime
	currentNumber, currentTime := utils.GetLatestBlockNumberAndBlockTime(client)
	fmt.Println("CurrentBlockNumber:", currentNumber, "CurrentBlockTime:", currentTime)

	//Get LatestBlock
	lastestBlock := utils.GetLatestBlock(client)
	header_Block := lastestBlock.Header()
	fmt.Println("Block Header:", utils.StructOrMapToJsonString(header_Block))
}
