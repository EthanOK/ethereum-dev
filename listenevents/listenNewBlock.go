package listenevents

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gocode.ethan/ethereum-dev/utils"
)

func ListenNewBlock(client *ethclient.Client) {
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Println(err)
		case header := <-headers:
			// fmt.Println(utils.StructOrMapToJsonString(header))

			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Block Hash:", block.Hash().Hex())
			fmt.Println("Block Number:", block.Number().Uint64())
			fmt.Println("Block Time:", block.Time())
			fmt.Println("BaseFee:", block.BaseFee())
			fmt.Println("Transactions Count:", len(block.Transactions()))

			for i, tx := range block.Transactions() {
				sender, err := client.TransactionSender(context.Background(), tx, block.Hash(), uint(i))
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Transaction Hash:", tx.Hash().Hex())
				fmt.Println("Sender Address:", sender.Hex())
				if tx.To() == nil {
					// 计算合约
					contractAddress := crypto.CreateAddress(sender, tx.Nonce())

					fmt.Println("To Address:", contractAddress.Hex(), "Contract Created")
				} else {
					fmt.Println("To Address:", tx.To().Hex())
				}

				fmt.Println("Value:", utils.WeiToEther(tx.Value()), "Ether")
				fmt.Println("Gas Limit:", tx.Gas())
				fmt.Println("Gas Price:", utils.WeiToGWei(tx.GasPrice()), "GWei")

				fmt.Println("Txn Type:", tx.Type())
				fmt.Println("Nonce:", tx.Nonce())
				fmt.Println("Position In Block:", i)
				fmt.Println("Input Data:", utils.Bytes2HexHas0xPrefix(tx.Data()))

			}
		}
	}
}
