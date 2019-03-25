package main

import (
	"fmt"
	"time"
)

var channel = make(chan int)

//定义一台打印机
func Printer(str string) {
	for _, ch := range str {
		fmt.Printf("%c", ch)
		time.Sleep(time.Second)
	}
}

func Person1() {
	Printer("hello")
	channel <- 1 //一直等待打印函数在获得的时间轮片中完成打印才能写入管道数据
}

func Person2() {
	<-channel //读不到阻塞
	Printer(" world")
}

//解决办法 1 加锁   2，channel
func main() {
	go Person1()
	go Person2()

	for {

	}
}
