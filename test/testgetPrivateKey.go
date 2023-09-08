package test

import (
	"fmt"

	"gocode.ethan/ethereum-dev/utils"
)

func TestPrivateKey() {
	privateKey, account := utils.GeneratePrivateKeyAndAccount()
	fmt.Println("privateKey:", privateKey)
	fmt.Println("account:", account)
}
