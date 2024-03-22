package controllers

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gocode.ethan/ethereum-dev/models"
	"gocode.ethan/ethereum-dev/utils"
)

func HandleERC6551AccountCreatedEvent(log types.Log, timestamp uint64) {

	// 解析事件参数
	contract := log.Address.Hex()
	blockNumber := log.BlockNumber
	txHash := log.TxHash.Hex()

	impl := common.BytesToAddress(log.Topics[1].Bytes()).Hex()
	tokenContract := common.BytesToAddress(log.Topics[2].Bytes()).Hex()
	tokenId := log.Topics[3].Big().String()

	address_, _ := abi.NewType("address", "", nil)
	bytes32_, _ := abi.NewType("bytes32", "", nil)
	uint256_, _ := abi.NewType("uint256", "", nil)

	arguments := abi.Arguments{
		{Type: address_},
		{Type: bytes32_},
		{Type: uint256_},
	}
	data := log.Data
	parsed, err := arguments.UnpackValues(data)
	if err != nil {
		panic(err)

	}

	tokenBoundAccount := parsed[0].(common.Address).Hex()
	salt_ := parsed[1].([32]byte)
	salt := common.BytesToHash(salt_[:]).Hex()

	chainId := parsed[2].(*big.Int).Uint64()

	accountCreated := models.ERC6551AccountCreated{
		ERC6551RegistryContract: contract,
		Timestamp:               timestamp,
		BlockNumber:             blockNumber,
		TxHash:                  txHash,
		Implementation:          impl,
		TokenContract:           tokenContract,
		TokenId:                 tokenId,
		TokenBoundAccount:       tokenBoundAccount,
		Salt:                    salt,
		ChainId:                 chainId}

	error_ := utils.GormDB_ERC6551.Create(&accountCreated).Error

	if error_ != nil {
		fmt.Println("插入数据重复")
	}
}
