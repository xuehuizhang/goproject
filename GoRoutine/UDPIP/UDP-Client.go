package main

import (
	"fmt"
	"net"
)

func main() {
	//创建用于通信的socket
	conn, err := net.Dial("udp", "127.0.0.1:8004")
	if err != nil {
		fmt.Println("net.Dial err", err)
		return
	}
	defer conn.Close()

	for i := 0; i < 1000000; i++ {
		//向服务器发送数据
		_, err = conn.Write([]byte("nihao"))
		if err != nil {
			fmt.Println("conn.Write err", err)
			return
		}

		buf := make([]byte, 4096)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err", err)
			return
		}
		fmt.Println("读取到服务器发送的数据", string(buf[:n]))
	}
}
