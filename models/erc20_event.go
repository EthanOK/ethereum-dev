package models

import "gorm.io/gorm"

type ERC20Transfer struct {
	gorm.Model
	Token       string
	From        string
	To          string
	Value       string
	TxHash      string `gorm:"type:varchar(100)"`
	Index       uint64
	Timestamp   uint64
	BlockNumber uint64
}

func (msg *ERC20Transfer) TableName() string {
	return "erc20_transfer"
}
