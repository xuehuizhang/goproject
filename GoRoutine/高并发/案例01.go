package main

import (
	"fmt"
	"time"
)

func sing() {
	for i := 0; i < 5; i++ {
		fmt.Println("唱隔壁泰山")
		time.Sleep(time.Second)
	}
}

func dance() {
	for i := 0; i < 5; i++ {
		fmt.Println("跳街舞")
		time.Sleep(time.Second)
	}
}

func main() {
	go sing()
	go dance()
	for {

	}
}
