package controllers

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gocode.ethan/ethereum-dev/models"
	"gocode.ethan/ethereum-dev/utils"
)

func HandleERC721TransferEvent(log types.Log, timestamp uint64) {
	// 处理ERC721转账事件

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

	err := utils.GormDB_EthereumDev.Create(&erc721Transfer).Error
	if err != nil {
		fmt.Println("Error inserting ERC20Transfer:", err)
	}

}
