package main

import "fmt"

func main()  {
	//iota 一般应用在枚举中
	//iota自增
	//在自增出现常量，后面的值会和常量相等
	const (
		one,two=iota+1,iota+2
		three,four
		fie,six
	)
	fmt.Println(one,two,three,four,fie,six)
}
