package utils

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"os"
	"regexp"

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

func GetAccount(privateKey *ecdsa.PrivateKey) string {
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return address.String()
}

func GetAccountByPublicKey(publicKeyECDSA *ecdsa.PublicKey) string {
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return address.String()
}

func GetAccountByPublicKeyBytes(pub []byte) string {
	pubkey, _ := crypto.UnmarshalPubkey(pub)
	return crypto.PubkeyToAddress(*pubkey).Hex()
}
func GetAccountByPublicKeyHex(publicKeyHex string) string {
	pub := hexutil.MustDecode(publicKeyHex)
	pubkey, _ := crypto.UnmarshalPubkey(pub)
	return crypto.PubkeyToAddress(*pubkey).Hex()
}

func GeneratePrivateKeyAndAccount() (string, string) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)

	}
	privateKeyBytes := crypto.FromECDSA(privateKey)

	// 无 0x前缀
	privateKeyHex := common.Bytes2Hex(privateKeyBytes)

	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	account := crypto.PubkeyToAddress(*publicKeyECDSA).String()
	return privateKeyHex, account

}

// str = "^0x0000[0-9a-fA-F]{36}$"
func GetPerfectAddress(str string) {
	fmt.Println("GetPerfectAddress:")
	re := regexp.MustCompile(str)
	result := false

	for !result {
		privateKey, account := GeneratePrivateKeyAndAccount()
		result = re.MatchString(account)
		if result {
			fmt.Println("privateKey:", privateKey)
			fmt.Println("account:", account)
		}
	}

}
