package test

import (
	"fmt"

	"gocode.ethan/ethereum-dev/utils"
)

func TestJsonTOString() {
	toStruct()
	toMap()
}
func toStruct() {
	type Person struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		City  string `json:"city"`
		Email string `json:"email"`
	}
	// jsonData := "{\"name\":\"Alice\",\"age\":30,\"city\":\"NewYork\",\"email\":\"alice@example.com\"}"

	jsonData_ := `{
    "name": "Alice",
    "age": 30,
    "city": "NewYork",
    "email": "alice@example.com"
	}`
	// 解析 JSON 数据到结构体
	var person Person
	utils.ParseJsonToStruct(jsonData_, &person)

	// 访问解析后的数据
	// fmt.Println("Person:", person)

	// 结构体 转 字符串
	jsonString := utils.StructOrMapToJsonString(person)

	// 打印 JSON 字符串
	fmt.Println(jsonString)
}

func toMap() {
	jsonData2_ := `{
		"jsonrpc": "2.0",
		"method": "eth_getLogs",
		"params": [
			{
				"fromBlock": "8888888",
				"topics": [
					"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
				]
			}
		],
		"id": 1
	}`
	var mapData map[string]interface{}
	utils.ParseJson(jsonData2_, &mapData)

	// 处理map数据
	// 获取 "params" 字段的值  转为切片
	params, ok := mapData["params"].([]interface{})
	if !ok {
		fmt.Println("无法获取 params 字段")
		return
	}
	// 单个元素params[0]  转为 map
	param, ok := params[0].(map[string]interface{})
	if !ok {
		fmt.Println("无法获取 params 中的第一个元素")
		return
	}
	fromBlock := param["fromBlock"]
	fmt.Println("fromBlock:", fromBlock)
	topics := param["topics"].([]interface{})
	topic := topics[0]
	fmt.Println("topics:", topic)

	jsonString2 := utils.StructOrMapToJsonString(mapData)

	// 打印 JSON 字符串
	fmt.Println(jsonString2)
}
