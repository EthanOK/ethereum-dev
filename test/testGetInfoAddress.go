package test

import (
	"fmt"

	"gocode.ethan/ethereum-dev/config"
	"gocode.ethan/ethereum-dev/getclient"
	"gocode.ethan/ethereum-dev/utils"
)

func TestGetInfoAddress() {

	address := "0x9a1f783ea8e236c44859a179e597791b7c066e84"
	fmt.Println("Un Format Address:", address)
	addressETH := utils.FormattedAddress(address)
	fmt.Println("Formatted Address:", addressETH)

	// Address isContract
	client, _ := getclient.GetEthClient()

	isContract := utils.CheckAddressIsContract(client, config.ACCOUNT_Vitalik)
	fmt.Println(config.ACCOUNT_Vitalik, "is Contract:", isContract)

	isContract_ := utils.CheckAddressIsContract(client, config.Token_DAI)
	fmt.Println(config.Token_DAI, "is Contract:", isContract_)

}
