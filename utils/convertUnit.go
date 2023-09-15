package utils

import (
	"math"
	"math/big"

	"github.com/shopspring/decimal"
)

// Wei To Ether
func WeiToEther0(weiBalance *big.Int) *big.Float {
	fbalance := new(big.Float)
	fbalance.SetString(weiBalance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	return ethValue
}

// Wei To Ether 可能最后几位会四舍五入
func WeiToEther(stringOrBigint interface{}) string {
	_decimal := ToDecimal(stringOrBigint, 18)
	return _decimal.String()
}

// Wei To GWei
func WeiToGWei(stringOrBigint interface{}) string {
	_decimal := ToDecimal(stringOrBigint, 9)
	return _decimal.String()
}

// Big To Decimal
func BigToDecimals(weiBalance *big.Int, decimals uint8) *big.Float {
	fbalance := new(big.Float)
	fbalance.SetString(weiBalance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(int(decimals))))
	return ethValue
}

// Ether To Wei
func EtherToWei(ethValue string) *big.Int {
	weiInt := ToWei(ethValue, 18)
	return weiInt
}

// ("0.1", 18) => 1e17
// ToWei: decimals to wei
func ToWei(iamount interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := iamount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}

// ToDecimal: wei to decimals
func ToDecimal(ivalue interface{}, decimals int) decimal.Decimal {
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)

	return result
}
