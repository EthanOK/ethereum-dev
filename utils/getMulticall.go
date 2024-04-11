package utils

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gocode.ethan/ethereum-dev/config"
	"gocode.ethan/ethereum-dev/contracts/multicall3"
)

func GetEthBalanceByMulticall3(client *ethclient.Client, account common.Address) *big.Int {
	// 创建一个Multicall3
	multicall3Address := common.HexToAddress(config.Multicall3)
	multicall3, err := multicall3.NewMulticall3(multicall3Address, client)
	if err != nil {
		return nil
	}
	// 调用Multicall3的 GetEthBalance 方法
	balance, err := multicall3.GetEthBalance(nil, account)
	if err != nil {
		return nil
	}
	return balance
}

type AggregateCall struct {
	Target   common.Address
	CallData []byte
}

func AggregateRead(client *ethclient.Client, caller common.Address, calls []AggregateCall) [][]byte {

	// 解析 Multicall3 合约ABI
	multicall3Abi, err := abi.JSON(strings.NewReader(multicall3.Multicall3ABI))
	if err != nil {
		return nil
	}

	calldata, err := multicall3Abi.Pack("aggregate", calls)
	if err != nil {
		fmt.Println("Failed to pack data:", err)
		return nil
	}

	result, err := client.CallContract(context.Background(), ethereum.CallMsg{
		// From: caller,
		To:   &config.Multicall3Address,
		Data: calldata,
	}, nil)

	if err != nil {
		fmt.Println("Failed to call contract:", err)
		return nil

	}
	// uint256 blockNumber, bytes[] memory returnData
	// 解析返回值

	parsedResult, err := multicall3Abi.Unpack("aggregate", result)
	if err != nil {
		fmt.Println("Failed to unpack result:", err)
		return nil
	}
	// blockNumber := parsedResult[0].(*big.Int)
	// fmt.Println("Block number:", blockNumber)
	returnData := parsedResult[1].([][]byte)
	// fmt.Println("Return data:", returnData)
	return returnData
}
