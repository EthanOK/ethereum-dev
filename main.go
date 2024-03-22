package main

import (
	"sync"

	"gocode.ethan/ethereum-dev/test"
	"gocode.ethan/ethereum-dev/utils"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	utils.InitDB()

	// test.TestTransferETH()
	// test.TestJsonTOString()
	// test.TestAccessMysqlDB()
	// test.TestGetInfoAccount()
	// test.TestGetInfoBlock()
	// test.TestAsyncFunc()
	// test.TestPrivateKey()
	// test.TestGetInfoAddress()
	// test.TestKeyStore()
	// test.TestFilterTransferLogs()
	// test.TestListenTransfer()
	// test.TestTransferERC20()
	// test.TestListenNewBlock()
	// test.TestGetTransactionReceipt()
	// test.TestSignVerify()
	// async.AsyncF()

	// test.TestCallContact()

	go func() {
		defer wg.Done()
		test.TestMulTopicsEvent()

	}()

	wg.Wait()
}
