package getclient

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func GetEthClient() (*ethclient.Client, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 使用 os 包中的 Getenv 函数读取环境变量
	ALCHEMY_RPC_HTTP := os.Getenv("ALCHEMY_RPC_HTTP")
	client, err := ethclient.Dial(ALCHEMY_RPC_HTTP)
	if err != nil {
		// log.Fatal(err)
		log.Println("get ethclient failure")
		return nil, err
	}
	return client, nil
}
func GetEthClient_G() (*ethclient.Client, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 使用 os 包中的 Getenv 函数读取环境变量
	ALCHEMY_RPC_HTTP_G := os.Getenv("ALCHEMY_RPC_HTTP_G")
	client, err := ethclient.Dial(ALCHEMY_RPC_HTTP_G)
	if err != nil {
		// log.Fatal(err)
		log.Println("get ethclient failure")
		return nil, err
	}
	return client, nil
}

func GetBscClient() (*ethclient.Client, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 使用 os 包中的 Getenv 函数读取环境变量
	ANKR_RPC_HTTP := os.Getenv("ANKR_RPC_HTTP")
	client, err := ethclient.Dial(ANKR_RPC_HTTP)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return client, nil
}
func GetBscClient_T() (*ethclient.Client, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 使用 os 包中的 Getenv 函数读取环境变量
	ANKR_RPC_HTTP_T := os.Getenv("ANKR_RPC_HTTP_T")
	client, err := ethclient.Dial(ANKR_RPC_HTTP_T)
	if err != nil {
		// log.Fatal(err)
		log.Println("get ethclient failure")
		return nil, err
	}
	return client, nil
}
