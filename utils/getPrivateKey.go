package utils

import (
	"crypto/ecdsa"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
)

func GeneratePrivateKeyNoPrefix() string {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)

	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	// 有 0x前缀
	// privateKeyHex := hexutil.Encode(privateKeyBytes)
	// 无 0x前缀
	privateKeyHex := common.Bytes2Hex(privateKeyBytes)

	return privateKeyHex
}

func GeneratePrivateKeyHasPrefix() string {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)

	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	// 有 0x前缀
	privateKeyHex := hexutil.Encode(privateKeyBytes)
	// 无 0x前缀
	// privateKeyHex := common.Bytes2Hex(privateKeyBytes)

	return privateKeyHex
}

func PrivateKeyToAddress(privateKeyHex string) string {
	has0x := has0xPrefix(privateKeyHex)
	if has0x {
		privateKeyHex = privateKeyHex[2:]
	}
	privateKey, _ := crypto.HexToECDSA(privateKeyHex)
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	return address.String()
}

func GetLocalPrivateKey() *ecdsa.PrivateKey {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 使用 os 包中的 Getenv 函数读取环境变量
	PRIVATEKEY_TEST := os.Getenv("PRIVATEKEY_TEST")
	privateKey, err := crypto.HexToECDSA(PRIVATEKEY_TEST)
	if err != nil {
		log.Fatal(err)
	}
	return privateKey
}
