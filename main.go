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

	// 0x9a1f783ea8e236c44859a179e597791b7c066e84 => 0x9a1F783EA8E236C44859a179E597791b7c066E84
	// addressE := utils.FormattedAddress("0x9a1f783ea8e236c44859a179e597791b7c066e84")
	// fmt.Println(addressE)

	// isContract := utils.CheckAddressIsContract(client, "0xdac17f958d2ee523a2206206994597c13d831ec7")
	// fmt.Println(isContract)

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
	test.TestPrivateKey()

}
