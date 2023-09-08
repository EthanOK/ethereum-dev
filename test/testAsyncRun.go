package test

import (
	"fmt"
	"time"

	"gocode.ethan/ethereum-dev/utils"
)

func TestAsyncFunc() {
	fmt.Println("start")
	fmt.Println(utils.GetSystemTimeStamp())
	// 使用通道来等待两个goroutines完成
	done := make(chan bool)

	go func() {
		asy1()
		done <- true
	}()

	go func() {
		asy2()
		done <- true
	}()

	// 等待两个goroutines完成
	<-done
	<-done

	// 增加一些延迟，以等待异步操作完成
	time.Sleep(1 * time.Second)
	fmt.Println("end")
	fmt.Println(utils.GetSystemTimeStamp())

}

func asy1() {
	time.Sleep(1 * time.Second)
	fmt.Println("asy1")
	fmt.Println(utils.GetSystemTimeStamp())

}

func asy2() {
	time.Sleep(2 * time.Second)
	fmt.Println("asy2")
	fmt.Println(utils.GetSystemTimeStamp())

}
