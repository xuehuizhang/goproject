package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	quit := make(chan bool)
	fmt.Println("当前时间：", time.Now())
	myTicker := time.NewTicker(time.Second) //周期定时器
	i := 0
	go func() {
		for {
			i++
			nowTime := <-myTicker.C
			fmt.Println("当前时间：", nowTime)
			if i == 8 {
				quit <- true
				runtime.Goexit()
			}
		}
	}()
	<-quit
}
