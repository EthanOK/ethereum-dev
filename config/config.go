package config

import (
	"github.com/ethereum/go-ethereum/common"
)

const ETHTIMEPER = 12

const BSCTIMEPER = 6

const ZeroAddress = "0x0000000000000000000000000000000000000000"

const DataSourceName = "root:root@tcp(192.168.0.173:3306)/aggregator_ethan"

const DB_Name_Ethereum_Dev = "root:root@tcp(127.0.0.1:3306)/ethereum-dev?charset=utf8mb4&parseTime=True&loc=Local"

const (
	USDT            = "0xdac17f958d2ee523a2206206994597c13d831ec7"
	ACCOUNT_Vitalik = "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045"
	Token_DAI       = "0x6B175474E89094C44Da98b954EedeAC495271d0F"
	ERC721_Rarible  = "0x60F80121C31A0d46B5279700f9DF786054aa5eE5"
	YGIO_TBSC       = "0xb06DcE9ae21c3b9163cD933E40c9EE563366b783"
	Multicall3      = "0xcA11bde05977b3631167028862bE2a173976CA11"
)

const (
	Transfer_Topic0              = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	ERC6551AccountCreated_Topic0 = "0x79f19b3655ee38b1ce526556b7731a20c8f218fbda4a3990b6cc4172fdf88722"
)

const (
	F_StartBlockNumber = "filter_start_block_number"
)

var Multicall3Address = common.HexToAddress(Multicall3)
