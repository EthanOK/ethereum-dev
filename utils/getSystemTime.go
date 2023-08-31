package utils

import (
	"time"
)

func GetSystemTimeStamp() int64 {
	// 获取当前时间
	currentTime := time.Now()
	// 获取时间戳（秒级）
	timestamp := currentTime.Unix()
	return timestamp
}
