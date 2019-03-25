package main

import "fmt"

func main() {
	go func() {

	}()

	for i := 0; i < 5; i++ {
		fmt.Println("主go程序")
		if i == 2 {
			break
		}
	}
}
