package listenevents

import (
	"log"

	"gocode.ethan/ethereum-dev/getclient"
)

func StartListenEvents() {
	/* 	ethclient, err := getclient.GetEthClient()
	   	if err != nil {
	   		log.Fatal(err)
	   		return
	   	} */
	bscclient, err := getclient.GetBscClient()
	if err != nil {
		log.Fatal(err)
		return
	}
	done := make(chan bool)

	go func() {
		// ListenTransferERC20(client, "0xdac17f958d2ee523a2206206994597c13d831ec7")
		// ListenTransferERC721_ETH(ethclient)
		ListenTransferERC721_BSC(bscclient)
		done <- true
	}()

	// 等待goroutines完成
	<-done
}
