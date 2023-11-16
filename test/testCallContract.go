package test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gocode.ethan/ethereum-dev/getclient"
	"gocode.ethan/ethereum-dev/utils"
	"golang.org/x/crypto/sha3"
)

func TestCallContact() {
	// Replace these values with your token address and DEX router address
	tokenAddress := common.HexToAddress("0x2B6d9f563d7CB1398C113ABb6DE615F9cDa3F0aA")     //Some random token
	dexRouterAddress := common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D") //Uniswap

	// Connect to a local Ethereum node using ethclient.Dial
	client, err := getclient.GetEthClient()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Call the getAmountIn function on the DEX router contract
	amountIn, err := callGetAmountIn(client, dexRouterAddress, tokenAddress)
	if err != nil {
		log.Fatal(err)
	}

	// Print the result
	fmt.Println("Result of getAmountIn:", amountIn)

}
func callGetAmountIn(client *ethclient.Client, dexRouter, token common.Address) (*big.Int, error) {
	// Prepare input parameters for the getAmountIn function
	amountOut := big.NewInt(100) // Example output value

	// Pack the input parameters
	data := getFunctionSelector("getAmountIn(uint256,uint256,uint256)", amountOut, big.NewInt(1000), big.NewInt(1000))

	fmt.Println(utils.Bytes2HexHas0xPrefix(data))

	// Create a CallMsg to perform a static call (no state changes)
	msg := ethereum.CallMsg{
		To:   &dexRouter,
		Data: data,
	}

	// Call the contract function
	output, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	// Unpack the output values
	result := new(big.Int)
	result.SetBytes(output)
	return result, nil
}

func getFunctionSelector(signature string, params ...*big.Int) []byte {
	// Calculate the function selector by taking the first 4 bytes of the hash of the signature

	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte(signature))
	functionSelector := hash.Sum(nil)[:4]

	// Combine the function selector and the packed parameters
	for _, param := range params {
		functionSelector = append(functionSelector, common.LeftPadBytes(param.Bytes(), 32)...)
	}
	return functionSelector
}
