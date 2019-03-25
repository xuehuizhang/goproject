package main

import (
	"fmt"
	"net"
	"os"
)

func RecvFile(conn net.Conn, fileName string) {
	//创建本地文件
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create err", err)
		return
	}
	//读取文件内容
	buf := make([]byte, 4096)
	for {
		n, _ := conn.Read(buf)
		if n == 0 {
			fmt.Println("读取文件完成")
			return
		}
		f.Write(buf[:n])
	}
}

func main() {
	//创建套接字
	listener, err := net.Listen("tcp", "127.0.0.1:8008")
	if err != nil {
		fmt.Println("net.Listen err", err)
		return
	}
	//监听客户端连接socket
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept err", err)
		return
	}
	//读取文件名
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	fileName := string(buf[:n])

	//向发送端写 ok
	_, err = conn.Write([]byte("ok"))
	if err != nil {
		fmt.Println("conn.Write err", err)
		return
	}

	RecvFile(conn, fileName)
}
