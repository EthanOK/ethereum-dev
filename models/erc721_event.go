package models

import "gorm.io/gorm"

type ERC721Transfer struct {
	gorm.Model
	Token       string
	From        string
	To          string
	TokenId     string
	TxHash      string `gorm:"type:varchar(100)"`
	Index       uint64
	Timestamp   uint64
	BlockNumber uint64
}

func (msg *ERC721Transfer) TableName() string {
	return "erc721_transfer"
}
