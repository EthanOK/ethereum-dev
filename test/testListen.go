package test

import (
	"gocode.ethan/ethereum-dev/getclient"
	"gocode.ethan/ethereum-dev/listenevents"
)

func TestListenTransfer() {

	// 不是实时监听 是每几秒扫一次区块
	listenevents.StartListenEvents()
}

func TestListenNewBlock() {
	// WSS
	client, _ := getclient.GetEthClientWSS_G()
	listenevents.ListenNewBlock(client)

}
