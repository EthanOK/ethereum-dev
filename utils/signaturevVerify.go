package utils

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// eth_sign
func Eth_Sign(privateKey *ecdsa.PrivateKey, data []byte) string {

	hash := crypto.Keccak256Hash(data)

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	return hexutil.Encode(signature)

}

// return address
func Recover_Eth_Sign(data []byte, signatureHex string) string {
	signature := hexutil.MustDecode(signatureHex)

	hash := crypto.Keccak256Hash(data)

	recoveredPub, err := crypto.Ecrecover(hash.Bytes(), signature)

	if err != nil {
		log.Fatal(err)
	}

	return GetAccountByPublicKeyBytes(recoveredPub)

}
