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
	balance := utils.GetBalance(client, address)
	fmt.Println("Account balance:", balance, "WEI")
	// wei to eth

	// Get the latest known block
	block := utils.GetLatestBlock(client)

	fmt.Println("block Hash:", block.Hash())

	blockTime := block.Time()
	fmt.Println("Latest block Number:", block.Number().Uint64())
	fmt.Println("Latest block Time:", blockTime)
	// 打印 本地时间戳
	systemTime := utils.GetSystemTimeStamp()

	fmt.Println("System Time:", systemTime)
}
