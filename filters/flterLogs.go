package filters

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gocode.ethan/ethereum-dev/utils"
)

func TransferLogsERC20(client *ethclient.Client, tokenAddress string, fromBlock string, toBlock string) {

	contractAddress := common.HexToAddress(tokenAddress)
	transferEventTopic := common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
	filterQuery := ethereum.FilterQuery{Addresses: []common.Address{contractAddress},
		FromBlock: utils.StringToBig(fromBlock),
		ToBlock:   utils.StringToBig(toBlock),
		Topics: [][]common.Hash{
			{transferEventTopic},
		}}
	logs, err := client.FilterLogs(context.Background(), filterQuery)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("logs length:", len(logs))
	for _, log := range logs {
		/* 		// 结构体转json数据
		   		jsonData_, err := json.Marshal(log)
		   		if err != nil {
		   			fmt.Println("Error encoding JSON:", err)
		   			return
		   		}
		   		// JSON数据转换为字符串
		   		jsonString := string(jsonData_)
		   		// 打印 JSON 字符串
		   		fmt.Println(jsonString) */

		// contractAddress := log.Address.Hex()
		txHash := log.TxHash.Hex()
		topics := log.Topics
		from := common.BigToAddress(topics[1].Big()).Hex()

		to := common.BigToAddress(topics[2].Big()).Hex()

		value := common.BytesToHash(log.Data).Big()
		fmt.Println("txHash:", txHash)
		fmt.Println(from, "=>", to)
		fmt.Println("amount:", value)
		fmt.Println("````````````````````````````````````````````````````````````````")

	}
}
