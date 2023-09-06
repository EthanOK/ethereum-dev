package filters

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"gocode.ethan/ethereum-dev/config"
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
		utils.StructToString(log)

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
			utils.StructToString(log)
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

func TransferLogsERC721(client *ethclient.Client, chainId int, fromBlock string, toBlock string) uint64 {

	transferEventTopic := common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
	filterQuery := ethereum.FilterQuery{
		FromBlock: utils.StringToBig(fromBlock),
		ToBlock:   nil,
		Topics: [][]common.Hash{
			{transferEventTopic},
		}}
	logs, err := client.FilterLogs(context.Background(), filterQuery)
	if err != nil {
		fmt.Println("FilterLogs err:", err)
		fromBlock_ := utils.StringToBig(fromBlock)
		return fromBlock_.Uint64() - 1
	}

	fromBlock_ := utils.StringToBig(fromBlock)
	// TODO: handledataERC721Print handledataERC721
	// handledataERC721Print(logs, fromBlock_.Uint64())
	lastBlockNumber := handledataERC721(logs, fromBlock_.Uint64(), chainId)
	return lastBlockNumber

}
func handledataERC721Print(logs []types.Log, lastBlockNumber uint64, chainId int) uint64 {

	len_ := len(logs)
	if len_ > 0 {
		for _, log := range logs {

			// token := log.Address.Hex()
			txHash := log.TxHash.Hex()
			// blockNumber := utils.Uint64ToString(log.BlockNumber)
			topics := log.Topics
			if len(topics) == 4 {

				from := common.BigToAddress(topics[1].Big()).Hex()

				to := common.BigToAddress(topics[2].Big()).Hex()
				if from != config.ZeroAddress && to != config.ZeroAddress {
					fmt.Println("flter blockNumber:", log.BlockNumber, txHash, chainId)
					lastBlockNumber = log.BlockNumber
					// utils.StructToString(log)

					// tokenId := topics[3].Big()

					// fmt.Println("txHash:", txHash)
					// fmt.Println(from, "=>", to)
					// fmt.Println("token:", token)
					// fmt.Println("tokenId:", tokenId)
					// fmt.Println("```````")
				}
			}

		}
		// fmt.Println("```````````````````````````````````````````````")
		return uint64(lastBlockNumber)
	} else {
		return uint64(lastBlockNumber)
	}

}
func handledataERC721(logs []types.Log, lastBlockNumber uint64, chainId int) uint64 {

	len_ := len(logs)
	if len_ > 0 {
		// 获取 mysql db
		db := utils.GetMysqlDB()
		// 关闭整个程序之前执行db.Close()
		defer db.Close()

		for _, log := range logs {

			token := log.Address.Hex()
			txHash := log.TxHash.Hex()
			blockNumber := utils.Uint64ToString(log.BlockNumber)
			topics := log.Topics
			if len(topics) == 4 {

				from := common.BigToAddress(topics[1].Big()).Hex()

				to := common.BigToAddress(topics[2].Big()).Hex()
				if from != config.ZeroAddress && to != config.ZeroAddress {
					fmt.Println("flter blockNumber:", log.BlockNumber, txHash, "chainId:", chainId)
					lastBlockNumber = log.BlockNumber
					// utils.StructToString(log)

					tokenId := topics[3].Big().String()
					insertSql := "INSERT IGNORE INTO event_transfer_erc721 (chainId,token, tokenId,fromAddress, toAddress, blockNumber, transactionHash) VALUES(?, ?, ?, ?, ?, ?, ?);"
					utils.Insert(db, insertSql, chainId, token, tokenId, from, to, blockNumber, txHash)

				}
			}

		}
		// fmt.Println("```````````````````````````````````````````````")
		return uint64(lastBlockNumber)
	} else {
		return uint64(lastBlockNumber)
	}

}
