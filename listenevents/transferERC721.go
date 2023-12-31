package listenevents

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"gocode.ethan/ethereum-dev/config"
	"gocode.ethan/ethereum-dev/filters"
	"gocode.ethan/ethereum-dev/utils"
)

func ListenTransferERC721(client *ethclient.Client, chainId int) {
	var interval_time int
	if chainId == 1 {
		interval_time = config.ETHTIMEPER
	} else if chainId == 56 {
		interval_time = config.BSCTIMEPER
	}

	currentBlockNumber, blockTime := utils.GetLatestBlockNumberAndBlockTime(client)

	interval := utils.GetSystemTimeStamp() - int64(blockTime)

	fmt.Println("interval:", interval)

	time.Sleep(time.Duration(interval_time-int(interval)+1) * time.Second)

	startBlockNumber := currentBlockNumber

	go func() {

		for {
			// 在每次循环中执行操作
			fmt.Println("startBlockNumber:", startBlockNumber)
			lastBlockNumber := filters.TransferLogsERC721(client, chainId,
				utils.Uint64ToString(startBlockNumber), "")
			fmt.Println("lastBlockNumber:", lastBlockNumber)
			startBlockNumber = lastBlockNumber + 1
			// 添加一些延迟
			time.Sleep(time.Duration(interval_time) * time.Second)
		}
	}()
	// 阻塞主goroutine，以防止程序退出
	select {}
}
