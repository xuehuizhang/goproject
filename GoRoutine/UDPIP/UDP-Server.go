package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	//组织一个udp的地址结构
	srv, err := net.ResolveUDPAddr("udp", "127.0.0.1:8003")
	if err != nil {
		fmt.Println("net.ResolveUDPAddr err:", err)
		return
	}
	fmt.Println("组织服务器udp地址结构完成")
	//创建用于通信的socket
	conn, err := net.ListenUDP("udp", srv)
	if err != nil {
		fmt.Println("net.ListenUDP err", err)
		return
	}
	defer conn.Close()
	fmt.Println("创建服务器socket完成")
	buf := make([]byte, 4096)
	//返回三个值  n 读取到的字节数  addr 客户端地址  err 错误
	n, cltAddr, err := conn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println("conn.Read err", err)
		return
	}
	//模拟处理数据
	fmt.Println(string(buf[:n]))
	fmt.Println("客户端地址：", cltAddr)

	//回写数据给客户端
	daytime := time.Now().String()

	_, err = conn.WriteToUDP([]byte(daytime), cltAddr)
	if err != nil {
		fmt.Println("conn.WriteToUDP err", err)
		return
	}
}
