package filters

import (
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"gocode.ethan/ethereum-dev/controllers"
	"gorm.io/gorm"
)

func HandleERC20TransferEvent(log types.Log, timestamp uint64) {
	controllers.HandleERC20TransferEvent(log, timestamp)
}

func BatchHandleERC20TransferEvent(db *gorm.DB, logs []types.Log, timestamp uint64) error {
	return controllers.BatchHandleERC20TransferEvent(db, logs, timestamp)
}

func HandleERC721TransferEvent(log types.Log, timestamp uint64) {
	controllers.HandleERC721TransferEvent(log, timestamp)
}

func BatchHandleERC721TransferEvent(db *gorm.DB, logs []types.Log, timestamp uint64) error {
	return controllers.BatchHandleERC721TransferEvent(db, logs, timestamp)
}

func HandleERC1155TransferEvent(log types.Log, timestamp uint64) {
	// 处理ERC1155转账事件
	fmt.Println("ERC1155 Transfer Event:", log.Address)
}

func HandleERC6551AccountCreatedEvent(log types.Log, timestamp uint64) {
	// 处理ERC6551账户创建事件
	// fmt.Println("ERC6551 Account Created Event:", log.Address)
	controllers.HandleERC6551AccountCreatedEvent(log, timestamp)
}

func BatchHandleERC6551AccountCreatedEvent(db *gorm.DB, logs []types.Log, timestamp uint64) error {
	return controllers.BatchHandleERC6551AccountCreatedEvent(db, logs, timestamp)
}
