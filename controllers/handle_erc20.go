package controllers

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gocode.ethan/ethereum-dev/models"
	"gocode.ethan/ethereum-dev/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func HandleERC20TransferEvent(log types.Log, timestamp uint64) {

	erc20Transfer := ProcessERC20TransferEvent(log, timestamp)

	err := utils.GormDB_EthereumDev.Create(&erc20Transfer).Error
	if err != nil {
		fmt.Println("Error inserting ERC20Transfer:", err)
	}

}

func BatchHandleERC20TransferEvent(db *gorm.DB, logs []types.Log, timestamp uint64) error {
	var erc20Transfers []models.ERC20Transfer

	for _, log := range logs {
		erc20Transfer := ProcessERC20TransferEvent(log, timestamp)
		erc20Transfers = append(erc20Transfers, erc20Transfer)
	}

	if len(erc20Transfers) > 0 {
		err := db.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(erc20Transfers, 100).Error
		if err != nil {
			// 记录日志并返回错误
			fmt.Println("Error inserting batch ERC20Transfers:", err)
			return err
		}
	}
	return nil
}

func ProcessERC20TransferEvent(log types.Log, timestamp uint64) models.ERC20Transfer {
	from := common.BytesToAddress(log.Topics[1].Bytes()).Hex()
	to := common.BytesToAddress(log.Topics[2].Bytes()).Hex()
	value := new(big.Int).SetBytes(log.Data).String()

	return models.ERC20Transfer{
		Token:       log.Address.Hex(),
		From:        from,
		To:          to,
		Value:       value,
		Timestamp:   timestamp,
		BlockNumber: log.BlockNumber,
		TxHash:      log.TxHash.Hex(),
		Index:       uint64(log.Index),
	}
}
