package utils

import (
	"math"
	"math/big"
)

// Wei To Ether
func WeiToEther(weiBalance *big.Int) *big.Float {
	fbalance := new(big.Float)
	fbalance.SetString(weiBalance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	return ethValue
}

// Ether To Wei
func EtherToWei(ethValue *big.Float) *big.Int {
	weiValue := new(big.Float).Mul(ethValue, big.NewFloat(math.Pow10(18)))
	weiInt := new(big.Int)
	weiValue.Int(weiInt)
	return weiInt
}
