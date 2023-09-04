package listenevents

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"gocode.ethan/ethereum-dev/filters"
	"gocode.ethan/ethereum-dev/utils"
)

func ListenTransferERC20(client *ethclient.Client, tokenAddress string) {
	currentBlockNumber, blockTime := utils.GetNowBlockNumberAndBlockTime(client)

	interval := utils.GetSystemTimeStamp() - int64(blockTime)

	fmt.Println("interval:", interval)

	time.Sleep(time.Duration(12-interval+1) * time.Second)

	lastBlockNumber := currentBlockNumber

	go func() {

		for {
			// 在每次循环中执行操作
			filters.TransferLogsERC20(client, tokenAddress,
				utils.Uint64ToString(lastBlockNumber), "")
			lastBlockNumber++
			// 添加一些延迟
			time.Sleep(12 * time.Second)
		}
	}()
	// 阻塞主goroutine，以防止程序退出
	select {}
}
