package test

import (
	"fmt"

	"gocode.ethan/ethereum-dev/utils"
)

func TestPrivateKey() {
	privateKey, account := utils.GeneratePrivateKeyAndAccount()
	fmt.Println("privateKey:", privateKey)
	fmt.Println("account:", account)

	utils.GetPerfectAddress("^0x0000[0-9a-fA-F]{36}$")
}
