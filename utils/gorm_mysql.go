package utils

import (
	"gocode.ethan/ethereum-dev/config"
	"gocode.ethan/ethereum-dev/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GormDB_EthereumDev *gorm.DB

func InitDB() {
	GormDB_EthereumDev = GetGormDB_EthereumDev()

}

func GetGormDB_EthereumDev() *gorm.DB {

	db, err := gorm.Open(mysql.Open(config.DB_Name_Ethereum_Dev), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.ERC6551AccountCreated{})
	db.AutoMigrate(&models.ERC20Transfer{})
	db.AutoMigrate(&models.ERC721Transfer{})
	db.AutoMigrate(&models.Config{})

	return db

}
