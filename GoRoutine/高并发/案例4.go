package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("cccccccccccccc")
	runtime.Goexit() //退出当前go程
	fmt.Println("dddddddddddddddd")
}

func main() {
	go func() {
		fmt.Println("aaaaaaaaaaaaaa")
		test()
		defer fmt.Println("bbbbbbbbbbbbbb")
	}()

	for {

	}
}
