package main

import (
	"gocode.ethan/ethereum-dev/test"
)

func main() {

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
	// test.TestKeyStore()
	test.TestFilterTransferLogs()

}
