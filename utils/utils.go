package utils

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/common"
)

func has0xPrefix(str string) bool {
	return len(str) >= 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X')
}
func FormattedAddress(addressHex string) string {

	address := common.HexToAddress(addressHex)

	// 格式化后的地址
	return address.Hex()

}

// string to big int
func StringToBig(str string) *big.Int {
	if str == "" {
		return nil
	}
	res, _ := new(big.Int).SetString(str, 10)
	return res

}

// string to big.Float
func StringToBigFloat(str string) *big.Float {
	if str == "" {
		return nil
	}
	res, _ := new(big.Float).SetString(str)
	return res

}

// uint64 to string
func Uint64ToString(num uint64) string {
	return new(big.Int).SetUint64(num).String()
	// return strconv.FormatUint(num, 10)
}

// uint64 to big
func Uint64ToBig(num uint64) *big.Int {
	return new(big.Int).SetUint64(num)
}

// structToString
func StructToString(s interface{}) string {
	// 结构体转json数据
	jsonData_, err := json.Marshal(s)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return ""
	}
	// JSON数据转换为字符串
	jsonString := string(jsonData_)

	return jsonString
}

func StructOrMapToJsonString(s interface{}) string {
	// Struct Or Map转json数据
	jsonData_, err := json.Marshal(s)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return ""
	}
	// JSON数据转换为字符串
	jsonString := string(jsonData_)

	return jsonString
}

func ParseJsonToStruct(jsonStr string, structData interface{}) {

	// 解析 JSON 字符串到 Struct
	err := json.Unmarshal([]byte(jsonStr), &structData)
	if err != nil {
		fmt.Println("解析 JSON 失败:", err)

	}

}

func ParseJson(jsonStr string, data *map[string]interface{}) {
	// // 创建一个 map 来存储解析后的 JSON 数据
	// var data map[string]interface{}

	// 解析 JSON 字符串到 map
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Println("解析 JSON 失败:", err)

	}

}

func GetLastFile(pathdir string) string {
	// 确保pathdir是一个有效的路径
	if _, err := os.Stat(pathdir); err != nil {
		fmt.Println("Invalid path:", pathdir)
		return ""
	}

	// 使用filepath.Walk()函数遍历pathdir目录并获取文件名
	var lastFile string
	var maxDepth int
	filepath.Walk(pathdir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error walking directory:", pathdir)
			return nil
		}

		if info.IsDir() {
			if maxDepth == 0 {
				maxDepth = 1
			} else {
				return filepath.SkipDir
			}
		} else {
			if maxDepth == 1 {
				lastFile = path
			}
		}

		return nil
	})

	// 返回最后一个文件的名称
	return lastFile
}
