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

	// controllers.CreateConfig(config.F_StartBlockNumber, "5534762")

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

	test.TestMulticall3()

	go func() {
		defer wg.Done()
		test.TestMulTopicsEvent()

	}()

	wg.Wait()
}
