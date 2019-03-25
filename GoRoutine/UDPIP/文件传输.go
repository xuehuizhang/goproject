package main

import (
	"fmt"
	"os"
)

func main() {
	list := os.Args
	if len(list) != 2 {
		fmt.Println("格式为： go fun xxx.go 文件名")
	}

	//获取文件属性
	path := list[1]
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("os.Stat err", err)
		return
	}
	fmt.Println("文件名：", fileInfo.Name())
	fmt.Println("文件大小：", fileInfo.Size())
}
