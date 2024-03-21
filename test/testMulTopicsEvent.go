package test

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"gocode.ethan/ethereum-dev/config"
	"gocode.ethan/ethereum-dev/filters"
	"gocode.ethan/ethereum-dev/getclient"
)

func TestMulTopicsEvent() {
	// 测试多主题事件

	client, err := getclient.GetEthClient_S()
	if err != nil {

		log.Fatal(err)

	}

	block, err_ := client.BlockByNumber(context.Background(), big.NewInt(5530545))
	if err_ != nil {

		fmt.Println("block not found")
		log.Fatal(err_)
	}

	ERC6551AccountCreated_Topic0 := common.HexToHash(config.ERC6551AccountCreated_Topic0)
	Transfer_Topic0 := common.HexToHash(config.Transfer_Topic0)
	topic0s := []common.Hash{ERC6551AccountCreated_Topic0, Transfer_Topic0}

	filters.MulTopicsEvent(client, block.Hash(), topic0s)
}
