package models

import (
	"gorm.io/gorm"
)

type Config struct {
	gorm.Model
	MapKey   string `gorm:"type:varchar(100);uniqueIndex"`
	MapValue string
}

func (ms *Config) TableName() string {
	// 返回表名
	return "config"
}
