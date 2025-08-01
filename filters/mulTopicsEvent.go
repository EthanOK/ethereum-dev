package filters

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"gocode.ethan/ethereum-dev/config"
	"gocode.ethan/ethereum-dev/utils"
)

func MulTopicsEvent(client *ethclient.Client, blockHash common.Hash, timestamp uint64, topic0s []common.Hash) {

	filterQuery := ethereum.FilterQuery{
		BlockHash: (*common.Hash)(&blockHash),
		Topics: [][]common.Hash{
			topic0s[:], // 多个 topic0，无需每个 topic0 调用一次
		},
	}

	logs, err := client.FilterLogs(context.Background(), filterQuery)
	if err != nil {
		log.Fatal(err)
	}
	HandleLogs(logs, timestamp)

}

func HandleLogs(logs []types.Log, timestamp uint64) {
	var erc20Logs []types.Log
	var erc721Logs []types.Log
	var erc6551Logs []types.Log

	for _, log := range logs {

		switch log.Topics[0] {

		case common.HexToHash(config.Transfer_Topic0):
			if len(log.Topics) == 3 {
				// HandleERC20TransferEvent(log, timestamp)
				erc20Logs = append(erc20Logs, log)
			} else if len(log.Topics) == 4 {
				// HandleERC721TransferEvent(log, timestamp)
				erc721Logs = append(erc721Logs, log)
			}

		case common.HexToHash(config.ERC6551AccountCreated_Topic0):

			if len(log.Topics) == 4 {
				// HandleERC6551AccountCreatedEvent(log, timestamp)
				erc6551Logs = append(erc6551Logs, log)
			}

		default:
			// 处理其他事件
		}
	}

	// 开始数据库事务
	tx := utils.GormDB_EthereumDev.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if len(erc20Logs) > 0 {
		err := BatchHandleERC20TransferEvent(tx, erc20Logs, timestamp)
		if err != nil {
			tx.Rollback()
			return
		}
	}

	if len(erc721Logs) > 0 {
		err := BatchHandleERC721TransferEvent(tx, erc721Logs, timestamp)
		if err != nil {
			tx.Rollback()
			return
		}
	}

	if len(erc6551Logs) > 0 {
		err := BatchHandleERC6551AccountCreatedEvent(tx, erc6551Logs, timestamp)
		if err != nil {
			tx.Rollback()
			return
		}
	}

	// 提交事务
	err := tx.Commit().Error
	if err != nil {
		log.Fatal("提交事务失败:", err)
	}

}
