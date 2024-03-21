package filters

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"gocode.ethan/ethereum-dev/config"
)

func MulTopicsEvent(client *ethclient.Client, blockHash common.Hash, topic0s []common.Hash) {

	filterQuery := ethereum.FilterQuery{
		BlockHash: (*common.Hash)(&blockHash),
		Topics: [][]common.Hash{
			topic0s[:],
		},
	}

	logs, err := client.FilterLogs(context.Background(), filterQuery)
	if err != nil {
		log.Fatal(err)
	}
	HandleLogs(logs)

}

func HandleLogs(logs []types.Log) {

	for _, log := range logs {

		// 处理 event
		switch log.Topics[0] {

		case common.HexToHash(config.Transfer_Topic0):
			if len(log.Topics) == 3 {
				HandleERC20TransferEvent(log)
			} else if len(log.Topics) == 4 {
				HandleERC721TransferEvent(log)
			}

		case common.HexToHash(config.ERC6551AccountCreated_Topic0):
			HandleERC6551AccountCreatedEvent(log)

		default:
			// 处理其他事件
		}

	}

}
