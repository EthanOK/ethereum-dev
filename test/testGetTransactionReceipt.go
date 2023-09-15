package test

import (
	"fmt"
	"math/big"

	"gocode.ethan/ethereum-dev/getclient"
	"gocode.ethan/ethereum-dev/utils"
)

func TestGetTransactionReceipt() {
	client, _ := getclient.GetEthClientWSS_G()
	txHex := "0xdfa6baaf0ebd33fec82721226136e8ff77f7162a5774373ed03509c98302273b"
	txReceipt := utils.GetTransactionReceipt(client, txHex)
	if txReceipt.Status == 0 {
		fmt.Println("The Transaction Failure")
	} else {
		fmt.Println("The Transaction Success")
	}

	fmt.Println(utils.StructOrMapToJsonString(txReceipt))
	// CumulativeGasUsed
	// 表示某个交易及其之前的所有交易在一个区块中的执行过程中所消耗的总燃料。
	// 这包括了当前交易的燃料消耗以及之前的所有交易的燃料消耗。
	fmt.Println("Position In Block:", txReceipt.TransactionIndex)
	fmt.Println("Block CumulativeGasUsed:", txReceipt.CumulativeGasUsed)
	fmt.Println("GasUsed:", txReceipt.GasUsed)
	fmt.Println("Gas Price:", txReceipt.EffectiveGasPrice)
	transactionFee := new(big.Int).Mul(txReceipt.EffectiveGasPrice, utils.Uint64ToBig(txReceipt.GasUsed))
	fmt.Println("Transaction Fee:", transactionFee, "Wei", utils.WeiToEther0(transactionFee), "Ether")
}
