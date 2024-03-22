package filters

import (
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"gocode.ethan/ethereum-dev/controllers"
)

func HandleERC20TransferEvent(log types.Log, timestamp uint64) {
	// 处理ERC20转账事件
	// fmt.Println("ERC20 Transfer Event:", log.Address)

}

func HandleERC721TransferEvent(log types.Log, timestamp uint64) {
	// 处理ERC721转账事件
	fmt.Println("ERC721 Transfer Event:", log.Address)

}

func HandleERC1155TransferEvent(log types.Log, timestamp uint64) {
	// 处理ERC1155转账事件
	fmt.Println("ERC1155 Transfer Event:", log.Address)
}

func HandleERC6551AccountCreatedEvent(log types.Log, timestamp uint64) {
	// 处理ERC6551账户创建事件
	fmt.Println("ERC6551 Account Created Event:", log.Address)
	controllers.HandleERC6551AccountCreatedEvent(log, timestamp)

}
