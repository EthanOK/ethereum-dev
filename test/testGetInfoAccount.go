package test

import (
	"fmt"

	"gocode.ethan/ethereum-dev/getclient"
	"gocode.ethan/ethereum-dev/utils"
)

const ACCOUNT_Vitalik = "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045"
const Token_DAI = "0x6B175474E89094C44Da98b954EedeAC495271d0F"

func TestGetInfoAccount() {
	client, _ := getclient.GetEthClient()

	// Get Account ETH Balance
	balance := utils.GetETHBalance(client, ACCOUNT_Vitalik)
	ethBalance := utils.WeiToEther(balance)
	fmt.Println("Vitalik Account balance:", ethBalance, "ETH")

	daiBalance, decimals := utils.GetERC20Balance(client, Token_DAI, ACCOUNT_Vitalik)

	fmt.Println("Vitalik Account DAI balance:", utils.BigToDecimals(daiBalance, decimals), "DAI")
}
