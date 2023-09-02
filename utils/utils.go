package utils

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func has0xPrefix(str string) bool {
	return len(str) >= 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X')
}
func FormattedAddress(addressHex string) string {

	address := common.HexToAddress(addressHex)

	// 格式化后的地址
	return address.Hex()

}

// string to big
func StringToBig(str string) *big.Int {
	if str == "" {
		return nil
	}
	res, _ := new(big.Int).SetString(str, 10)
	return res

}

// uint64 to string
func Uint64ToString(num uint64) string {
	return new(big.Int).SetUint64(num).String()
	// return strconv.FormatUint(num, 10)
}

// uint64 to big
func Uint64ToBig(num uint64) *big.Int {
	return new(big.Int).SetUint64(num)
}
