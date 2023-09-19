package test

import (
	"fmt"

	"gocode.ethan/ethereum-dev/utils"
)

func TestSignVerify() {
	fmt.Println("eth_sign")
	eth_sign()
	fmt.Println("person_sign")
	person_sign()
}

func eth_sign() {
	privateKey := utils.GetLocalPrivateKey()

	signAddress := utils.GetAccount(privateKey)

	message := "hello"

	signatureHex := utils.Eth_Sign(message, privateKey)

	fmt.Println("signatureHex:", signatureHex)

	recoverAddress := utils.Recover_Eth_Sign(message, signatureHex)

	fmt.Println(signAddress, recoverAddress)

	fmt.Println("Verify Result:", signAddress == recoverAddress)
	fmt.Println("`````````````````````````````````````````````````")
}

func person_sign() {
	privateKey := utils.GetLocalPrivateKey()

	signAddress := utils.GetAccount(privateKey)

	// text
	text := "hello"
	signatureHex := utils.PersonalSign(text, privateKey)

	fmt.Println("signatureHex:", signatureHex)

	recoverAddress := utils.Recover_PersonalSign(text, signatureHex)

	fmt.Println(signAddress, recoverAddress)

	fmt.Println("Verify Result:", signAddress == recoverAddress)
	fmt.Println("`````````````````````````````````````````````````")

	// hex []byte
	bytesData := utils.HexToBytes("0xf6896007477ab25a659f87c4f8c5e3baac32547bf305e77aa57743046e10578b")
	signatureHex_ := utils.PersonalSign(
		bytesData, privateKey)
	fmt.Println("signatureHex_:", signatureHex_)
	recoverAddress_ := utils.Recover_PersonalSign(bytesData, signatureHex_)

	fmt.Println(signAddress, recoverAddress_)

	fmt.Println("Verify Result:", signAddress == recoverAddress_)
	fmt.Println("`````````````````````````````````````````````````")

}
