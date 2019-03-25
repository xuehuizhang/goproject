package main

import "fmt"

func main() {
	ch := make(chan string)
	fmt.Println(len(ch), cap(ch)) //无缓冲的channel  len与cap始终为0
	go func() {
		for i := 0; i < 2; i++ {
			fmt.Println("子go程打印=", i)
		}
		ch <- "子go程打印完毕"
	}()
	str := <-ch
	fmt.Println(str)
}
