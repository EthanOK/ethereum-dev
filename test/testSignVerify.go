package test

import (
	"fmt"

	"gocode.ethan/ethereum-dev/utils"
)

func TestSignVerify() {
	eth_sign()
}

func eth_sign() {
	privateKey := utils.GetLocalPrivateKey()

	signAddress := utils.GetAccount(privateKey)

	data := []byte("hello")

	signatureHex := utils.Eth_Sign(privateKey, data)

	fmt.Println("signatureHex:", signatureHex)

	recoverAddress := utils.Recover_Eth_Sign(data, signatureHex)

	fmt.Println(signAddress, recoverAddress)

	fmt.Println("Verify Result:", signAddress == recoverAddress)
}
func person_sign() {}
