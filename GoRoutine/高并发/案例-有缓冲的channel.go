package main

import (
	"fmt"
)

//有缓冲的channel
func main() {
	ch := make(chan int, 3)
	fmt.Println("cap=", cap(ch), "len=", len(ch))
	go func() {
		for i := 0; i < 8; i++ {
			ch <- i
			fmt.Println("子go写入：i=", i, "cap=", cap(ch), "len=", len(ch))
		}
	}()
	//time.Sleep(time.Second*3)
	for i := 0; i < 8; i++ {
		n := <-ch
		fmt.Println("主go读取：", n)
	}
}
