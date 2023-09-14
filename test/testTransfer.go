package test

import (
	"gocode.ethan/ethereum-dev/config"
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

func TestTransferERC20() {
	client, _ := getclient.GetBscClient_T()
	privateKey := utils.GetLocalPrivateKey()
	amount := utils.EtherToWei("2").String()
	toAddress := "0x53188E798f2657576c9de8905478F46ac2f24b67"
	utils.TransferERC20(client, privateKey, config.YGIO_TBSC,
		toAddress, amount)
}
