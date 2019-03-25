package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int) //数据通信的chan
	quit := make(chan bool)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
		quit <- true
	}()
	for {
		select {
		case num := <-ch:
			fmt.Println("读到：", num)
			time.Sleep(time.Second)
		case <-quit:
			return
		}
	}
}
