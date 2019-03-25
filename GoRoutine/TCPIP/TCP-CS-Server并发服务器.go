package main

import (
	"fmt"
	"net"
	"strings"
)

func main()  {
	//创建监听套接字
	//listener,err:=net.Listen("tcp","127.0.0.1:8000")
	listener,err:=net.Listen("tcp","172.18.2.34:8000")
	if err!=nil{
		fmt.Println("Net.Listen err",err)
		return
	}
	defer listener.Close()

	//监听客户端请求
	for{
		fmt.Println("正在监听客户端连接...")
		conn,err:=listener.Accept()
		if err!=nil{
			fmt.Println("listener.Accept err:",err)
			return
		}
		//具体完成服务器和客户端的通信
		go HandlerConnect(conn)
	}
}

func HandlerConnect(conn net.Conn)  {
	defer conn.Close()
	//读取客户端网络地址
	addr:=conn.RemoteAddr()
	fmt.Println(addr,"客户端成功连接")
	//循环读取客户端数据
	buf:=make([]byte,4096)
	for{
		n,err:=conn.Read(buf)
		if "exit\r\n"==string(buf[:n])||"exit\n"==string(buf[:n]){
			fmt.Println(addr,"服务器接收到客户端退出请求 退出")
			return
		}
		if n==0{
			fmt.Println(addr,"服务器检测到客户端关闭...")
			return
		}
		if err!=nil{
			fmt.Println("conn.Read err:",err)
			return
		}

		fmt.Println("服务器读到数据:",string(buf[:n]))
		//小写转大写
		str:=strings.ToUpper(string(buf[:n]))
		conn.Write([]byte(str))
	}
}