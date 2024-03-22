package utils

import (
	"gocode.ethan/ethereum-dev/config"
	"gocode.ethan/ethereum-dev/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GormDB_ERC6551 *gorm.DB

func InitDB() {
	GormDB_ERC6551 = GetGormDB_ERC6551()

}

func GetGormDB_ERC6551() *gorm.DB {

	db, err := gorm.Open(mysql.Open(config.DataSourceName_ERC6551), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&models.ERC6551AccountCreated{})

	return db

}
