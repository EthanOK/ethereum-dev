package controllers

import (
	"fmt"

	"gocode.ethan/ethereum-dev/models"
	"gocode.ethan/ethereum-dev/utils"
)

func CreateConfig(key string, value string) {
	config := models.Config{MapKey: key, MapValue: value}

	err := utils.GormDB_EthereumDev.Create(&config).Error

	if err != nil {
		fmt.Println("err: CreateConfig")
	}
}

func GetConfigValue(key string) string {

	var config models.Config

	err := utils.GormDB_EthereumDev.Where("map_key = ?", key).First(&config).Error
	if err != nil {
		fmt.Println("err: GetConfigValue")
	}

	return config.MapValue
}

func UpdataConfig(key string, value string) {
	err := utils.GormDB_EthereumDev.Model(&models.Config{}).
		Where("map_key = ?", key).Update("map_value", value).Error

	if err != nil {
		fmt.Println("err: UpdataConfig")
	}
}
