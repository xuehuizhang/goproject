package main

import (
	"fmt"
	"runtime"
)

func main() {

	go func() {
		for {
			fmt.Println("this is goroutine test")
		}
	}()

	for {
		runtime.Gosched()
		fmt.Println("this is main test")
	}
}
