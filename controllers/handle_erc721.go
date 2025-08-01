package controllers

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gocode.ethan/ethereum-dev/models"
	"gocode.ethan/ethereum-dev/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func HandleERC721TransferEvent(log types.Log, timestamp uint64) {

	erc721Transfer := ProcessERC721TransferEvent(log, timestamp)

	err := utils.GormDB_EthereumDev.Create(&erc721Transfer).Error
	if err != nil {
		fmt.Println("Error inserting ERC20Transfer:", err)
	}

}

func BatchHandleERC721TransferEvent(db *gorm.DB, logs []types.Log, timestamp uint64) error {
	var erc721Transfers []models.ERC721Transfer

	for _, log := range logs {
		erc721Transfer := ProcessERC721TransferEvent(log, timestamp)
		erc721Transfers = append(erc721Transfers, erc721Transfer)
	}

	if len(erc721Transfers) > 0 {
		err := db.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(erc721Transfers, 100).Error
		if err != nil {
			// 记录日志并返回错误
			fmt.Println("Error inserting batch ERC721Transfers:", err)
			return err
		}
	}
	return nil
}

func ProcessERC721TransferEvent(log types.Log, timestamp uint64) models.ERC721Transfer {

	from := common.BytesToAddress(log.Topics[1].Bytes()).Hex()
	to := common.BytesToAddress(log.Topics[2].Bytes()).Hex()
	tokenId := log.Topics[3].Big().String()

	erc721Transfer := models.ERC721Transfer{
		Token:       log.Address.Hex(),
		From:        from,
		To:          to,
		TokenId:     tokenId,
		Timestamp:   timestamp,
		BlockNumber: log.BlockNumber,
		TxHash:      log.TxHash.Hex(),
		Index:       uint64(log.Index),
	}
	return erc721Transfer
}
