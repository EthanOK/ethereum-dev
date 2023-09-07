package main

import (
	"fmt"

	"gocode.ethan/ethereum-dev/getclient"
	"gocode.ethan/ethereum-dev/test"
	"gocode.ethan/ethereum-dev/utils"
)

var password = "goodboy"

func main() {
	client, err := getclient.GetEthClient()
	if err != nil {

		// log.Fatal(err)
		// return
	}
	currentBlockNumber, _ := utils.GetNowBlockNumberAndBlockTime(client)
	fmt.Println("current Block Number:", currentBlockNumber)
	nonce := utils.GetAccountNonce(client, "0x95B3Aad7f20E78a0e4DcEB9c23beca4e55ebdDF6")
	fmt.Println("current Nonce:", nonce)
	/*
		   	account := "0x95B3Aad7f20E78a0e4DcEB9c23beca4e55ebdDF6"
		   	// Get the balance of an account
		   	balance := utils.GetETHBalance(client, account)
		   	// fmt.Println("Account balance:", balance, "WEI")
		   	fmt.Println("Account balance:", utils.WeiToEther(balance), "ETH")
		   	// Get the latest known block
		   	blockNumber, blockTime := utils.GetNowBlockNumberAndBlockTime(client)
		   	fmt.Println("Latest Block Number:", blockNumber)
		   	fmt.Println("Latest Block Time:", blockTime)
		   	// Get the balance of USDT
		   	contractUSDT := "0xdAC17F958D2ee523a2206206994597C13D831ec7"
		   	metadata := utils.GetERC20TokenMetadata(client, contractUSDT)
		   	symbol := metadata.Symbol
		   	decimals := metadata.Decimals
		   	usdtBalance := utils.GetERC20Balance(client, contractUSDT, account)
		   	fmt.Println("Account balance:", utils.BigToDecimals(usdtBalance,decimals), symbol)


		privateKey := utils.GeneratePrivateKeyHasPrefix()
		fmt.Println("privateKey:", privateKey)
		address := utils.PrivateKeyToAddress(privateKey)
		isAddress := utils.CheckIsAddress(address)
		fmt.Println(address, "is Address:", isAddress)

		type Person struct {
			Name  string `json:"name"`
			Age   int    `json:"age"`
			City  string `json:"city"`
			Email string `json:"email"`
		}
		jsonData := "{\"name\":\"Alice\",\"age\":30,\"city\":\"NewYork\",\"email\":\"alice@example.com\"}"
		//jsonData := `{"name":"Alice","age":30,"city":"NewYork","email":"alice@example.com"}`

		// 解析 JSON 数据到结构体
		var person Person
		err_ := json.Unmarshal([]byte(jsonData), &person)
		if err_ != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}

		// 访问解析后的数据
		fmt.Println("Name:", person.Name)
		fmt.Println("Age:", person.Age)
		fmt.Println("City:", person.City)
		fmt.Println("Email:", person.Email)
		// 使用 json.Marshal 将结构体转换为 JSON 字符串
		jsonData_, err := json.Marshal(person)
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}

		// 将 JSON 数据转换为字符串
		jsonString := string(jsonData_)

		// 打印 JSON 字符串
		fmt.Println(jsonString)
	*/

	// KeyStore
	/* pathdir := "./temp"
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

	/* db := utils.GetMysqlDB(
	// 关闭整个程序之前执行db.Close()
	defer db.Close()

	utils.Query(db, "SELECT id, `key`, value FROM aggregator_ethan.`system`;")

	utils.Insert(db, "INSERT INTO `system` (`key`, value) VALUES(?, ?);", "ens", "ether.eth")

	utils.Query(db, "SELECT id, `key`, value FROM aggregator_ethan.`system`;") */

	/* 	go listenevents.StartListenEvents()

	   	select {} */

	test.TransferETH()
}
