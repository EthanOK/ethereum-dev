package filters

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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
	fmt.Println("blockNumber:", fromBlock)
	// handledataERC20(logs)
	handledataERC20Print(logs)

}
func handledataERC20(logs []types.Log) {
	// 获取 mysql db
	db := utils.GetMysqlDB()
	// 关闭整个程序之前执行db.Close()
	defer db.Close()

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

		token := log.Address.Hex()
		txHash := log.TxHash.Hex()
		blockNumber := utils.Uint64ToString(log.BlockNumber)
		topics := log.Topics
		from := common.BigToAddress(topics[1].Big()).Hex()

		to := common.BigToAddress(topics[2].Big()).Hex()

		value := common.BytesToHash(log.Data).Big().String()

		// fmt.Println("txHash:", txHash)
		// fmt.Println(from, "=>", to)
		// fmt.Println("amount:", value)
		// fmt.Println("````````````````````````````````````````````````````````````````")
		insertSql := "INSERT IGNORE INTO event_transfer_erc20 (token, fromAddress, toAddress, value, blockNumber, transactionHash) VALUES(?, ?, ?, ?, ?, ?);"
		utils.Insert(db, insertSql, token, from, to, value, blockNumber, txHash)
	}
}

func handledataERC20Print(logs []types.Log) {
	if len(logs) > 0 {
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

			// token := log.Address.Hex()
			txHash := log.TxHash.Hex()
			// blockNumber := utils.Uint64ToString(log.BlockNumber)
			topics := log.Topics
			from := common.BigToAddress(topics[1].Big()).Hex()

			to := common.BigToAddress(topics[2].Big()).Hex()

			value := common.BytesToHash(log.Data).Big().String()

			fmt.Println("txHash:", txHash)
			fmt.Println(from, "=>", to)
			fmt.Println("amount:", value)
			fmt.Println("```````")

		}
		fmt.Println("```````````````````````````````````````````````")
	}

}

func TransferLogsERC721(client *ethclient.Client, fromBlock string, toBlock string) {

	transferEventTopic := common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
	filterQuery := ethereum.FilterQuery{
		FromBlock: utils.StringToBig(fromBlock),
		ToBlock:   utils.StringToBig(toBlock),
		Topics: [][]common.Hash{
			{transferEventTopic},
		}}
	logs, err := client.FilterLogs(context.Background(), filterQuery)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("blockNumber:", fromBlock)
	// handledataERC721(logs)
	handledataERC721Print(logs)

}
func handledataERC721Print(logs []types.Log) {
	if len(logs) > 0 {
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

			token := log.Address.Hex()
			txHash := log.TxHash.Hex()
			// blockNumber := utils.Uint64ToString(log.BlockNumber)
			topics := log.Topics
			if len(topics) == 4 {
				from := common.BigToAddress(topics[1].Big()).Hex()

				to := common.BigToAddress(topics[2].Big()).Hex()

				tokenId := topics[3].Big()

				fmt.Println("txHash:", txHash)
				fmt.Println(from, "=>", to)
				fmt.Println("token:", token)
				fmt.Println("tokenId:", tokenId)
				fmt.Println("```````")
			}

		}
		fmt.Println("```````````````````````````````````````````````")
	}

}
