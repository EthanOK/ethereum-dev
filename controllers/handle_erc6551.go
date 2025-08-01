package controllers

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gocode.ethan/ethereum-dev/contracts/erc6551"
	"gocode.ethan/ethereum-dev/models"
	"gocode.ethan/ethereum-dev/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func HandleERC6551AccountCreatedEvent(log types.Log, timestamp uint64) {

	accountCreated := ProcessERC6551AccountCreatedEvent(log, timestamp)

	error_ := utils.GormDB_EthereumDev.Create(&accountCreated).Error

	if error_ != nil {
		fmt.Println("插入数据重复")
	}
}

func BatchHandleERC6551AccountCreatedEvent(db *gorm.DB, logs []types.Log, timestamp uint64) error {
	var accountCreateds []models.ERC6551AccountCreated
	for _, log := range logs {
		accountCreated := ProcessERC6551AccountCreatedEvent(log, timestamp)
		accountCreateds = append(accountCreateds, accountCreated)
	}

	if len(accountCreateds) > 0 {
		err := db.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(accountCreateds, 100).Error
		if err != nil {
			// 记录日志并返回错误
			fmt.Println("Error inserting batch ERC721Transfers:", err)
			return err
		}
	}
	return nil

}

func ProcessERC6551AccountCreatedEvent(log types.Log, timestamp uint64) models.ERC6551AccountCreated {

	impl := common.BytesToAddress(log.Topics[1].Bytes()).Hex()
	tokenContract := common.BytesToAddress(log.Topics[2].Bytes()).Hex()
	tokenId := log.Topics[3].Big().String()

	erc6551RegistryAbi, _ := abi.JSON(strings.NewReader(erc6551.ERC6551RegistryABI))

	parsed, _ := erc6551RegistryAbi.Unpack("ERC6551AccountCreated", log.Data)

	tokenBoundAccount := parsed[0].(common.Address).Hex()
	salt_ := parsed[1].([32]byte)
	salt := common.BytesToHash(salt_[:]).Hex()

	chainId := parsed[2].(*big.Int).Uint64()

	accountCreated := models.ERC6551AccountCreated{
		ERC6551RegistryContract: log.Address.Hex(),
		Timestamp:               timestamp,
		BlockNumber:             log.BlockNumber,
		TxHash:                  log.TxHash.Hex(),
		Implementation:          impl,
		TokenContract:           tokenContract,
		TokenId:                 tokenId,
		TokenBoundAccount:       tokenBoundAccount,
		Salt:                    salt,
		ChainId:                 chainId}
	return accountCreated
}
