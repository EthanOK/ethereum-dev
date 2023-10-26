package utils

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// eth_sign
func Eth_Sign(data interface{}, privateKey *ecdsa.PrivateKey) string {
	var fullMessage string
	if str, ok := data.(string); ok {
		// data 是一个字符串
		fullMessage = fmt.Sprintf("%s", str)
	} else if bytesData, ok := data.([]byte); ok {
		// data 是一个字节数组
		fullMessage = fmt.Sprintf("%s", bytesData)
	} else {
		// data 既不是字符串也不是字节数组，返回自定义错误
		log.Fatal("data type error")
	}

	hash := crypto.Keccak256Hash([]byte(fullMessage))

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	return hexutil.Encode(signature)

}

// return address
func Recover_Eth_Sign(data interface{}, signatureHex string) string {
	signature := hexutil.MustDecode(signatureHex)

	var fullMessage string
	if str, ok := data.(string); ok {
		// data 是一个字符串
		fullMessage = fmt.Sprintf("%s", str)
	} else if bytesData, ok := data.([]byte); ok {
		// data 是一个字节数组
		fullMessage = fmt.Sprintf("%s", bytesData)
	} else {
		// data 既不是字符串也不是字节数组，返回自定义错误
		log.Fatal("data type error")
	}

	hash := crypto.Keccak256Hash([]byte(fullMessage))

	recoveredPub, err := crypto.Ecrecover(hash.Bytes(), signature)

	if err != nil {
		log.Fatal(err)
	}

	return GetAccountByPublicKeyBytes(recoveredPub)

}

func Recover_PersonalSignText(message string, signatureHex string) string {
	signature := hexutil.MustDecode(signatureHex)

	signature[64] -= 27

	fullMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)

	hash := crypto.Keccak256Hash([]byte(fullMessage))

	recoveredPub, err := crypto.Ecrecover(hash.Bytes(), signature)

	if err != nil {
		log.Fatal(err)
	}

	return GetAccountByPublicKeyBytes(recoveredPub)

}

// 恢复签名，返回恢复地址
func Recover_PersonalSign(data interface{}, signatureHex string) string {
	signature := hexutil.MustDecode(signatureHex)

	signature[64] -= 27

	var fullMessage string
	if str, ok := data.(string); ok {
		// data 是一个字符串
		fullMessage = fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(str), str)
	} else if bytesData, ok := data.([]byte); ok {
		// data 是一个字节数组
		fullMessage = fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(bytesData), bytesData)
	} else {
		// data 既不是字符串也不是字节数组，返回自定义错误
		log.Fatal("data type error")
	}

	hash := crypto.Keccak256Hash([]byte(fullMessage))

	recoveredPub, err := crypto.Ecrecover(hash.Bytes(), signature)

	if err != nil {
		log.Fatal(err)
	}

	return GetAccountByPublicKeyBytes(recoveredPub)

}

// 签名文本消息
func PersonalSignText(message string, privateKey *ecdsa.PrivateKey) (string, error) {
	fullMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)

	hash := crypto.Keccak256Hash([]byte(fullMessage))

	signatureBytes, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return "", err
	}
	signatureBytes[64] += 27
	return hexutil.Encode(signatureBytes), nil
}

// 签名二进制 []byte
func PersonalSignBytes(data []byte, privateKey *ecdsa.PrivateKey) (string, error) {
	fullMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)

	hash := crypto.Keccak256Hash([]byte(fullMessage))

	signatureBytes, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return "", err
	}
	signatureBytes[64] += 27
	return hexutil.Encode(signatureBytes), nil
}

// data Type: string or []byte; Return signatureHex
func PersonalSign(data interface{}, privateKey *ecdsa.PrivateKey) string {
	var fullMessage string
	if str, ok := data.(string); ok {
		// data 是一个字符串
		fullMessage = fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(str), str)
	} else if bytesData, ok := data.([]byte); ok {
		// data 是一个字节数组
		fullMessage = fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(bytesData), bytesData)
	} else {
		// data 既不是字符串也不是字节数组，返回自定义错误
		log.Fatal("data type error")
	}

	hash := crypto.Keccak256Hash([]byte(fullMessage))

	signatureBytes, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal("crypto.Sign error")
	}
	signatureBytes[64] += 27
	return hexutil.Encode(signatureBytes)
}

func HexToBytes(hex string) []byte {
	return common.FromHex(hex)
}
