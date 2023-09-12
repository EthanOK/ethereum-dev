package test

import (
	"gocode.ethan/ethereum-dev/utils"
)

func TestKeyStore() {
	pathdir := "./temp"
	var password = "goodboy"
	utils.GetKeyStore(pathdir, password)
	// 获取pathdir路径下的最后一个文件路径
	kspath := utils.GetLastFile(pathdir)
	// 导入密钥文件
	utils.ImportKeyStore(kspath, password)

	utils.GetAccountsByKeyStoreInWalletsFile()

}
