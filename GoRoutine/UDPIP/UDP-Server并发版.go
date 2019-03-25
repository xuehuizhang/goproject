package main

import (
	"fmt"
	"net"
	"time"
)

//udp并发并不需要起go程
func main() {
	//创建udp地址结构
	svrAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8004")
	if err != nil {
		fmt.Println("net.ResolveUDPAddr err")
		return
	}
	fmt.Println("创建udp地址结构完成")
	//创建udpsocket
	conn, err := net.ListenUDP("udp", svrAddr)
	if err != nil {
		fmt.Println("net.ListenUDP err", err)
		return
	}
	fmt.Println("服务端Socket创建完成")

	buf := make([]byte, 4096)
	for {
		n, cliAdr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("conn.ReadFromUDP", err)
			return
		}
		fmt.Println(cliAdr, "从客户端读取到数据:", buf[:n])
		go func() {
			//往客户端写数据
			dayTime := time.Now().String()

			_, err = conn.WriteToUDP([]byte(dayTime), cliAdr)
			if err != nil {
				fmt.Println("conn.WriteToUDP err", err)
			}
		}()
	}
}
