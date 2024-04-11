package test

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"gocode.ethan/ethereum-dev/config"
	"gocode.ethan/ethereum-dev/getclient"
	"gocode.ethan/ethereum-dev/utils"
)

func TestMulticall3() {

	client, err := getclient.GetEthClient_S()
	if err != nil {
		return
	}
	accountAdress := common.HexToAddress(config.ACCOUNT_Vitalik)
	eth_balance := utils.GetEthBalanceByMulticall3(client, accountAdress)
	fmt.Println("eth_balance:", eth_balance)

	var params []utils.AggregateCall
	// symbol(): 0x95d89b41; totalSupply(): 0x18160ddd; name(): 0x06fdde03
	params = append(params, utils.AggregateCall{
		Target:   common.HexToAddress("0x779877A7B0D9E8603169DdbD7836e478b4624789"),
		CallData: common.FromHex("0x18160ddd")})

	params = append(params, utils.AggregateCall{
		Target:   common.HexToAddress("0x709B78B36b7208f668A3823c1d1992C0805E4f4d"),
		CallData: common.FromHex("0x95d89b41")})

	returnData := utils.AggregateRead(client, common.HexToAddress(config.ACCOUNT_Vitalik), params)

	totalSupply := utils.Bytes2Big(returnData[0])

	symbol := utils.DynamicBytes2String(returnData[1])
	fmt.Println("totalSupply:", totalSupply)
	fmt.Println("symbol:", symbol)

}
