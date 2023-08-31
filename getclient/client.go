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
	ALCHEMY_PRC_HTTP := os.Getenv("ALCHEMY_PRC_HTTP")
	client, err := ethclient.Dial(ALCHEMY_PRC_HTTP)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return client, nil
}
