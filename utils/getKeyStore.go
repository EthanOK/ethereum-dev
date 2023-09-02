package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func GetKeyStore(pathdir string, password string) {
	ks := keystore.NewKeyStore(pathdir, keystore.StandardScryptN, keystore.StandardScryptP)

	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(account.Address.Hex())
}

func ImportKeyStore(keystoreFilePath string, password string) {
	// 创建keystore实例
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)

	// 读取keystore文件内容
	jsonBytes, err := os.ReadFile(keystoreFilePath)
	if err != nil {
		log.Fatal(err)
	}
	// 导入keystore文件，解锁账户
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal(err)
	}

	// 打印导入的账户地址
	fmt.Println("导入的账户地址:", account.Address.Hex())

}

func GetAccountsByKeyStoreInWalletsFile() {
	// 创建keystore实例
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	//定义 string 数组
	var str []string

	accounts := ks.Accounts()
	for _, account := range accounts {

		// 向 str [] 添加元素
		str = append(str, account.Address.Hex())
	}
	fmt.Println(str)

}
