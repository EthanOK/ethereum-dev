package listenevents

import (
	"log"

	"gocode.ethan/ethereum-dev/getclient"
)

func StartListenEvents() {

	done := make(chan bool)

	/* 	go func() {
		ethclient, err := getclient.GetEthClient()
		if err != nil {
			log.Fatal(err)
			return
		}
		// ListenTransferERC20(client, "0xdac17f958d2ee523a2206206994597c13d831ec7")
		ListenTransferERC721_ETH(ethclient, 1)

		done <- true

	}() */

	go func() {
		bscclient, err := getclient.GetBscClient()
		if err != nil {
			log.Fatal(err)
			return
		}
		// ListenTransferERC20(client, "0xdac17f958d2ee523a2206206994597c13d831ec7")
		// ListenTransferERC721_ETH(ethclient, 1)
		ListenTransferERC721_ETH(bscclient, 56)
		done <- true

	}()

	// 等待goroutines完成
	<-done
	// <-done

}
