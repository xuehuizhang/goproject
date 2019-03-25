package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func SendFile(conn net.Conn, filePath string) {
	//从本地读取文件
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("os.Open err", err)
		return
	}
	buf := make([]byte, 4096)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("读取文件完成")
			} else {
				fmt.Println("读取文件报错")
			}
			return
		}
		conn.Write(buf[:n])
	}

}

func main() {
	//读取文件名
	list := os.Args
	if len(list) != 2 {
		fmt.Println("格式为 go run xxx.go 文件名")
		return
	}
	filePath := list[1]
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("os.Stat err", err)
		return
	}
	fileName := fileInfo.Name()

	//建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:8008")
	if err != nil {
		fmt.Println("net.Dial err", err)
		return
	}
	//发送文件名
	_, err = conn.Write([]byte(fileName))
	if err != nil {
		fmt.Println("conn.Write err", err)
		return
	}
	//读取服务端回执
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err", err)
		return
	}
	if "ok" == string(buf[:n]) {
		SendFile(conn, filePath)
	} else {
		fmt.Println("服务端拒绝接收文件")
	}
}
