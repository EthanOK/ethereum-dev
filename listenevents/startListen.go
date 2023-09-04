package listenevents

import (
	"log"

	"gocode.ethan/ethereum-dev/getclient"
)

func StartListenEvents() {
	client, err := getclient.GetEthClient()
	if err != nil {
		log.Fatal(err)
		return
	}
	done := make(chan bool)

	go func() {
		ListenTransferERC20(client, "0xdac17f958d2ee523a2206206994597c13d831ec7")
		done <- true
	}()

	// 等待goroutines完成
	<-done
}
