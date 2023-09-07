package test

import (
	"gocode.ethan/ethereum-dev/getclient"
	"gocode.ethan/ethereum-dev/utils"
)

func TransferETH() {
	client, _ := getclient.GetBscClient_T()
	privateKey := utils.GetLocalPrivateKey()
	toAddress := ""
	amount := utils.StringToBig("1000000")

	utils.TransferETH(client, privateKey, toAddress, amount)
}
