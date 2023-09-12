package test

import (
	"gocode.ethan/ethereum-dev/config"
	"gocode.ethan/ethereum-dev/filters"
	"gocode.ethan/ethereum-dev/getclient"
	"gocode.ethan/ethereum-dev/utils"
)

func TestFilterTransferLogs() {
	client, _ := getclient.GetEthClient()
	blockNumber := utils.GetLatestBlockNumber(client)

	filters.TransferLogsERC20(client, config.USDT,
		utils.Uint64ToString(blockNumber.Uint64()-1), "")
}
