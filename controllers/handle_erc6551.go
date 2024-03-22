package controllers

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gocode.ethan/ethereum-dev/models"
	"gocode.ethan/ethereum-dev/token/erc6551"
	"gocode.ethan/ethereum-dev/utils"
)

func HandleERC6551AccountCreatedEvent(log types.Log, timestamp uint64) {

	// 解析 Event 参数
	contract := log.Address.Hex()
	blockNumber := log.BlockNumber
	txHash := log.TxHash.Hex()

	// 解析 Topics 参数
	impl := common.BytesToAddress(log.Topics[1].Bytes()).Hex()
	tokenContract := common.BytesToAddress(log.Topics[2].Bytes()).Hex()
	tokenId := log.Topics[3].Big().String()

	// TODO：方式一 解析 log.Data 参数
	// address_, _ := abi.NewType("address", "", nil)
	// bytes32_, _ := abi.NewType("bytes32", "", nil)
	// uint256_, _ := abi.NewType("uint256", "", nil)

	// arguments := abi.Arguments{
	// 	{Type: address_},
	// 	{Type: bytes32_},
	// 	{Type: uint256_},
	// }

	// parsed, err := arguments.UnpackValues(log.Data)
	// if err != nil {
	// 	panic(err)
	// }

	// TODO：方式二 解析 log.Data 参数
	erc6551RegistryAbi, _ := abi.JSON(strings.NewReader(erc6551.ERC6551RegistryABI))

	parsed, _ := erc6551RegistryAbi.Unpack("ERC6551AccountCreated", log.Data)

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
