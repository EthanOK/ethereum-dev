package test

import (
	"fmt"

	"gocode.ethan/ethereum-dev/getclient"
	"gocode.ethan/ethereum-dev/utils"
)

const (
	ACCOUNT_Vitalik = "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045"
	Token_DAI       = "0x6B175474E89094C44Da98b954EedeAC495271d0F"
	ERC721_Rarible  = "0x60F80121C31A0d46B5279700f9DF786054aa5eE5"
)

func TestGetInfoAccount() {
	// Get Client
	client, _ := getclient.GetEthClient()

	// Check Is Address
	isAddress := utils.CheckIsAddress(ACCOUNT_Vitalik)
	fmt.Println(ACCOUNT_Vitalik, "is Address:", isAddress)

	// Get Account ETH Balance
	balance := utils.GetETHBalance(client, ACCOUNT_Vitalik)
	ethBalance := utils.WeiToEther(balance)
	fmt.Println("Vitalik Account balance:", ethBalance, "ETH")

	// Get Account ERC20 Balance
	daiBalance, decimals := utils.GetERC20Balance(client, Token_DAI, ACCOUNT_Vitalik)
	fmt.Println("Vitalik Account ERC20:DAI balance:", utils.BigToDecimals(daiBalance, decimals), "DAI")

	// Get Account ERC721 Balance
	balance_R, symbol := utils.GetERC721Balance(client, ERC721_Rarible, ACCOUNT_Vitalik)
	fmt.Println("Vitalik Account ERC721:Rarible balance:", balance_R, symbol)

	// Get Account Nonce
	nonce := utils.GetAccountNonce(client, ACCOUNT_Vitalik)
	fmt.Println("Vitalik Account Nonce:", nonce)
}
