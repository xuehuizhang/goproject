package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU())
	//fmt.Println(runtime.CPUProfile())
	n := runtime.GOMAXPROCS(2)
	fmt.Println(n)
	for {
		go fmt.Print(0) //子
		fmt.Print(1)    //主
	}
}
