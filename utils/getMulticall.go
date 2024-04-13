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

type AggregateResult struct {
	BlockNumber *big.Int
	ReturnData  [][]byte
}

func AggregateRead(client *ethclient.Client, calls []AggregateCall) AggregateResult {
	var result AggregateResult

	multicall3Abi, err := abi.JSON(strings.NewReader(multicall3.Multicall3ABI))
	if err != nil {
		return result
	}

	calldata, err := multicall3Abi.Pack("aggregate", calls)
	if err != nil {
		fmt.Println("Failed to pack data:", err)
		return result
	}

	result_, err := client.CallContract(context.Background(), ethereum.CallMsg{
		To:   &config.Multicall3Address,
		Data: calldata,
	}, nil)
	if err != nil {
		fmt.Println("Failed to call contract:", err)
		return result

	}

	parsedResult, err := multicall3Abi.Unpack("aggregate", result_)
	if err != nil {
		fmt.Println("Failed to unpack result:", err)
		return result
	}

	for _, v := range parsedResult {
		// fmt.Printf("%T\n", v)
		if blockNumber, ok := v.(*big.Int); ok {
			result.BlockNumber = blockNumber
		} else if returnData, ok := v.([][]uint8); ok {
			result.ReturnData = returnData
		}

	}

	return result
}

type AggregateCall3 struct {
	Target       common.Address
	AllowFailure bool
	CallData     []byte
}

type Aggregate3Result struct {
	Success    bool
	ReturnData []byte
}

func Aggregate3Read(client *ethclient.Client, calls []AggregateCall3) (results []Aggregate3Result) {

	multicall3Abi, err := abi.JSON(strings.NewReader(multicall3.Multicall3ABI))
	if err != nil {
		return nil
	}

	calldata, err := multicall3Abi.Pack("aggregate3", calls)
	if err != nil {
		fmt.Println("Failed to pack data:", err)
		return nil
	}

	result, err := client.CallContract(context.Background(), ethereum.CallMsg{
		To:   &config.Multicall3Address,
		Data: calldata,
	}, nil)
	if err != nil {
		fmt.Println("Failed to call contract:", err)
		return nil

	}

	parsedResult, err := multicall3Abi.Unpack("aggregate3", result)
	if err != nil {
		fmt.Println("Failed to unpack result:", err)
		return nil
	}

	for _, v := range parsedResult {
		// fmt.Printf("%T\n", v)
		for _, vv := range v.([]struct {
			Success    bool    "json:\"success\""
			ReturnData []uint8 "json:\"returnData\""
		}) {
			results = append(results, Aggregate3Result{
				Success:    vv.Success,
				ReturnData: vv.ReturnData,
			})
		}
	}

	return results
}
