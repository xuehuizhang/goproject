package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			fmt.Println("子go写入：", i)
		}
	}()
	time.Sleep(time.Second * 2)
	for i := 0; i < 5; i++ {
		n := <-ch
		fmt.Println("主go读取：", n)
	}
}
