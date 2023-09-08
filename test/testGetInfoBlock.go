package test

import (
	"fmt"

	"gocode.ethan/ethereum-dev/getclient"
	"gocode.ethan/ethereum-dev/utils"
)

func TestGetInfoBlock() {
	// Get Client
	client, _ := getclient.GetEthClient()

	// Get CurrentBlockNumber CurrentBlockTime
	currentNumber, currentTime := utils.GetLatestBlockNumberAndBlockTime(client)
	fmt.Println("CurrentBlockNumber:", currentNumber)
	fmt.Println("CurrentBlockTime:", currentTime)

	//Get LatestBlock
	lastestBlock := utils.GetLatestBlock(client)
	header_Block := lastestBlock.Header()
	fmt.Println("Block Header:", utils.StructOrMapToJsonString(header_Block))

	/*
			{
		    "parentHash": "0x99b98af891b78f3f47ae737d8e1819ed5f4e7106e2b8aeb5b4790d8ee09cd4be",
		    "sha3Uncles": "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347",
		    "miner": "0x690b9a9e9aa1c9db991c7721a92d351db4fac990",
		    "stateRoot": "0xe714a0a1b7e23e27088e4f9947c4fe5c1ebd98b55a88026a429f84b9e1d9834f",
		    "transactionsRoot": "0x7eff4ddf5c04bc8d0aeb96c3fc5e453a232a850fca19ded7a7fe151e3b43cde7",
		    "receiptsRoot": "0x61d3dd8f50b53fe933be06be37457f1eb6b22e177ec8259ea2fff3d0e20f99cf",
		    "logsBloom": "0x773531d6e7bf248d52767ef8ca534e421fba9934dca9ea9af8a98f6bed322c2ff03363cde24d43e4d0d9b66de6bff79d0fbd36b5c9c9be60bef75afde5f8f080eababb58f66ad8fffe474a7edc7a10bc9464551eeb7fbb0de770f7dfd5693e51368726e3dec47f738dcdd3578d002c2fe75bbc5a96dcfe97ee567cd1dada97248db607571b4fb3654dff522de6dba6eefc126c89ef65095f55bda7c566942ca83e7b93d6ebfa70c1ebeb8af57cb425e7479367b56f41ee7acfd0f7bf71f14b80d7b218d35a77bca4047aebb6cf7edbcb1af5b203dfeff154f8a13e0af355a12d2576a0fa81bb84bf6ec79aa1fa1f932ec7a89f9e5f08abd38de9dd9d02c95c79",
		    "difficulty": "0x0",
		    "number": "0x1140a75",
		    "gasLimit": "0x1c95111",
		    "gasUsed": "0x190ff83",
		    "timestamp": "0x64fae013",
		    "extraData": "0x406275696c64657230783639",
		    "mixHash": "0x4ea65d65bc3bc96aa69b248bfd7a94a9056822fcc6c9141bbd9dceb686d04b01",
		    "nonce": "0x0000000000000000",
		    "baseFeePerGas": "0x29ce54db2",
		    "withdrawalsRoot": "0x88a59fffdf50e8ea4180b631637924ab6f2dfd5660a774b98a27191ef41fc1de",
		    "blobGasUsed": null,
		    "excessBlobGas": null,
		    "hash": "0x87bc1286e6a15eecec6b6bca71cec96f764eda256f3c67155738ec1955df1952"
		}
	*/
}
