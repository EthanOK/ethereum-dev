package test

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gocode.ethan/ethereum-dev/config"
	"gocode.ethan/ethereum-dev/filters"
	"gocode.ethan/ethereum-dev/getclient"
)

var wg sync.WaitGroup

func TestMulTopicsEvent() {
	// 测试多主题事件
	wg.Add(1)

	client, err := getclient.GetEthClient_S()
	if err != nil {

		log.Fatal(err)

	}

	go startListenEvent(client, big.NewInt(5531180))

	wg.Wait()

}
func startListenEvent(client *ethclient.Client, startBlockNumber *big.Int) {
	defer wg.Done()

	for {

		block, err_ := client.BlockByNumber(context.Background(), startBlockNumber)
		if err_ != nil {

			fmt.Println("区块还没产生：", startBlockNumber)

		} else {
			ERC6551AccountCreated_Topic0 := common.HexToHash(config.ERC6551AccountCreated_Topic0)
			Transfer_Topic0 := common.HexToHash(config.Transfer_Topic0)
			topic0s := []common.Hash{ERC6551AccountCreated_Topic0, Transfer_Topic0}

			filters.MulTopicsEvent(client, block.Hash(), topic0s)
			fmt.Println("此区块已完成: ", startBlockNumber)
			// 获取下一个区块
			startBlockNumber.Add(startBlockNumber, big.NewInt(1))
		}

		time.Sleep(6 * time.Second)

	}

}
