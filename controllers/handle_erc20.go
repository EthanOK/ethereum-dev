package controllers

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"gocode.ethan/ethereum-dev/models"
	"gocode.ethan/ethereum-dev/utils"
)

func HandleERC20TransferEvent(log types.Log, timestamp uint64) {
	// 处理ERC20转账事件

	from := common.BytesToAddress(log.Topics[1].Bytes()).Hex()
	to := common.BytesToAddress(log.Topics[2].Bytes()).Hex()
	value := new(big.Int).SetBytes(log.Data).String()
	// hash(from, to, value)
	uniqueHash := crypto.Keccak256Hash(
		log.Address.Bytes(),
		log.Topics[1].Bytes(),
		log.Topics[2].Bytes(),
		log.Data, log.TxHash.Bytes()).Hex()

	erc20Transfer := models.ERC20Transfer{
		Token:       log.Address.Hex(),
		From:        from,
		To:          to,
		Value:       value,
		Timestamp:   timestamp,
		BlockNumber: log.BlockNumber,
		TxHash:      log.TxHash.Hex(),
		UniqueHash:  uniqueHash,
	}

	err := utils.GormDB_EthereumDev.Create(&erc20Transfer).Error
	if err != nil {
		fmt.Println("Error inserting ERC20Transfer:", err)
	}

}
