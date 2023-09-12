package main

import (
	"gocode.ethan/ethereum-dev/test"
)

func main() {

	// KeyStore
	/* pathdir := "./temp"
	var password = "goodboy"
	utils.GetKeyStore(pathdir, password)

	kspath := "./temp/UTC--2023-09-02T06-37-57.096776700Z--99e972b2c4689fba2aa8f31a6cebbdd12e3242c9"
	utils.ImportKeyStore(kspath, password)

	utils.GetAccountsByKeyStoreInWalletsFile()
	*/

	// filters.TransferLogsERC20(client, "0xdac17f958d2ee523a2206206994597c13d831ec7",
	// 	utils.Uint64ToString(currentBlockNumber), "")

	/* 	go async.AsyncF()

	   	// 使用通道来等待goroutines完成
	   	done := make(chan bool)

	   	go func() {
	   		listenevents.ListenTransferERC20(client, "0xdac17f958d2ee523a2206206994597c13d831ec7")
	   		done <- true
	   	}()

	   	// 等待goroutines完成
	   	<-done */

	/* 	go listenevents.StartListenEvents()

	   	select {} */

	// test.TestTransferETH()
	// test.TestJsonTOString()
	// test.TestAccessMysqlDB()
	// test.TestGetInfoAccount()
	// test.TestGetInfoBlock()
	// test.TestAsyncFunc()
	// test.TestPrivateKey()
	// test.TestGetInfoAddress()
	test.TestKeyStore()

}
