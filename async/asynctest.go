package async

import (
	"fmt"
	"time"
)

func AsyncF() bool {
	fmt.Println("start")
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
	return true
}

func asy1() {
	time.Sleep(1 * time.Second)
	fmt.Println("asy1")

}
func asy2() {
	time.Sleep(2 * time.Second)
	fmt.Println("asy2")

}
