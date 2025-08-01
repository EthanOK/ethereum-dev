package test

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gocode.ethan/ethereum-dev/config"
	"gocode.ethan/ethereum-dev/controllers"
	"gocode.ethan/ethereum-dev/filters"
	"gocode.ethan/ethereum-dev/getclient"
	"gocode.ethan/ethereum-dev/utils"
)

func TestMulTopicsEvent() {
	// 测试多主题事件

	client, err := getclient.GetEthClient_S()
	if err != nil {

		log.Fatal(err)

	}
	startBlockNumber := controllers.GetConfigValue(config.F_StartBlockNumber)

	startListenEvent(client, utils.StringToBig(startBlockNumber))

}

func startListenEvent(client *ethclient.Client, startBlockNumber *big.Int) {

	// 定时任务 1s 执行一次
	ticker := time.NewTicker(1 * time.Second)

	for {
		block, err_ := client.BlockByNumber(context.Background(), startBlockNumber)
		if err_ != nil {

			fmt.Println("区块还没产生：", startBlockNumber)

			if err_.Error() == "not found" {
				ticker.Reset(12 * time.Second)
			}

		} else {
			ERC6551AccountCreated_Topic0 := common.HexToHash(config.ERC6551AccountCreated_Topic0)
			Transfer_Topic0 := common.HexToHash(config.Transfer_Topic0)

			// TODO：将待查询的事件topic0写入topic0s
			topic0s := []common.Hash{ERC6551AccountCreated_Topic0, Transfer_Topic0}

			filters.MulTopicsEvent(client, block.Hash(), block.Time(), topic0s)
			fmt.Println("此区块已完成: ", startBlockNumber)

			// 保存下一个区块至数据库
			startBlockNumber.Add(startBlockNumber, big.NewInt(1))
			controllers.UpdataConfig(config.F_StartBlockNumber, startBlockNumber.String())
		}
		<-ticker.C
	}
}
