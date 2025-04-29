package controllers

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gocode.ethan/ethereum-dev/models"
	"gocode.ethan/ethereum-dev/utils"
)

func HandleERC20TransferEvent(log types.Log, timestamp uint64) {
	// 处理ERC20转账事件

	from := common.BytesToAddress(log.Topics[1].Bytes()).Hex()
	to := common.BytesToAddress(log.Topics[2].Bytes()).Hex()
	value := new(big.Int).SetBytes(log.Data).String()

	erc20Transfer := models.ERC20Transfer{
		Token:       log.Address.Hex(),
		From:        from,
		To:          to,
		Value:       value,
		Timestamp:   timestamp,
		BlockNumber: log.BlockNumber,
		TxHash:      log.TxHash.Hex(),
		Index:       uint64(log.Index),
	}

	err := utils.GormDB_EthereumDev.Create(&erc20Transfer).Error
	if err != nil {
		fmt.Println("Error inserting ERC20Transfer:", err)
	}

}
