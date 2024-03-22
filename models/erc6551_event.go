package models

import (
	"gorm.io/gorm"
)

type ERC6551AccountCreated struct {
	gorm.Model
	TokenContract           string
	TokenId                 string
	TokenBoundAccount       string `gorm:"type:varchar(100);uniqueIndex"`
	Implementation          string
	ERC6551RegistryContract string
	TxHash                  string
	Timestamp               uint64
	BlockNumber             uint64
	ChainId                 uint64
	Salt                    string
}

func (msg *ERC6551AccountCreated) TableName() string {
	return "erc6551_account_created"

}
