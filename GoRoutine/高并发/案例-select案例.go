package main

import (
	"fmt"
	"runtime"
)

//斐波那契额数列
// 1 1 2 3 5 8 13....
func main() {
	ch := make(chan int)
	quit := make(chan bool)

	go fib(ch, quit)
	x, y := 1, 1
	for i := 1; i < 20; i++ {
		ch <- x
		x, y = y, x+y
	}
	quit <- true
}

func fib(ch <-chan int, quit <-chan bool) {
	for {
		select {
		case num := <-ch:
			fmt.Print(num, " ")
		case <-quit:
			runtime.Goexit()
		}
	}
}
