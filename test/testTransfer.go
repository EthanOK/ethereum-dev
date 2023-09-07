package test

import (
	"gocode.ethan/ethereum-dev/getclient"
	"gocode.ethan/ethereum-dev/utils"
)

func TestTransferETH() {
	client, _ := getclient.GetBscClient_T()
	privateKey := utils.GetLocalPrivateKey()
	toAddress := "0x53188E798f2657576c9de8905478F46ac2f24b67"
	amount := "1000000"
	utils.TransferETH(client, privateKey, toAddress, amount)
}
