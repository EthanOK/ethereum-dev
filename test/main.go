package main

import (
	"encoding/json"
	"fmt"
	"log"

	"gocode.ethan/ethereum-dev/getclient"
	"gocode.ethan/ethereum-dev/utils"
)

func main() {

		client, err := getclient.GetEthClient()
	   	if err != nil {
	   		log.Fatal(err)
	   		return
	   	}
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

}
