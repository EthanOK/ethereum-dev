package utils

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
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
