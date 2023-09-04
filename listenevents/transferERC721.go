package listenevents

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"gocode.ethan/ethereum-dev/config"
	"gocode.ethan/ethereum-dev/filters"
	"gocode.ethan/ethereum-dev/utils"
)

func ListenTransferERC721_ETH(client *ethclient.Client) {
	interval_time := config.ETHTIMEPER
	currentBlockNumber, blockTime := utils.GetNowBlockNumberAndBlockTime(client)

	interval := utils.GetSystemTimeStamp() - int64(blockTime)

	fmt.Println("interval:", interval)

	time.Sleep(time.Duration(interval_time-int(interval)+1) * time.Second)

	lastBlockNumber := currentBlockNumber

	go func() {

		for {
			// 在每次循环中执行操作
			filters.TransferLogsERC721(client,
				utils.Uint64ToString(lastBlockNumber), "")
			lastBlockNumber++
			// 添加一些延迟
			time.Sleep(time.Duration(interval_time) * time.Second)
		}
	}()
	// 阻塞主goroutine，以防止程序退出
	select {}
}
func ListenTransferERC721_BSC(client *ethclient.Client) {
	interval_time := config.BSCTIMEPER
	currentBlockNumber, blockTime := utils.GetNowBlockNumberAndBlockTime(client)

	interval := utils.GetSystemTimeStamp() - int64(blockTime)

	fmt.Println("interval:", interval)

	time.Sleep(time.Duration(interval_time-int(interval)+1) * time.Second)

	startBlockNumber := currentBlockNumber

	go func() {

		for {
			fmt.Println("startBlockNumber:", startBlockNumber)
			// 在每次循环中执行操作
			lastBlockNumber := filters.TransferLogsERC721(client,
				utils.Uint64ToString(startBlockNumber), "")
			fmt.Println("lastBlockNumber", lastBlockNumber)
			fmt.Println("```````````````````````````````````````````````")
			startBlockNumber = lastBlockNumber + 1
			// 添加一些延迟
			time.Sleep(time.Duration(interval_time) * time.Second)
		}
	}()
	// 阻塞主goroutine，以防止程序退出
	select {}
}
