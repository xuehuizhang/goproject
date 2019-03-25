package main

import (
	"fmt"
	"net"
)

func main()  {
	listener,err:=net.Listen("tcp","127.0.0.1:8080") //创建一个用户监听的套接字
	if err!=nil{
		fmt.Println("net listen err:",err)
		return
	}
	defer listener.Close()
	//阻塞客户端链接请求
	fmt.Println("服务器等待客户端建立连接....")
	//成功建立连接 返回用户通信的socket
	conn,err:=listener.Accept()
	if err!=nil{
		fmt.Println("accept err:",err)
	}
	defer conn.Close()
	fmt.Println("成功建立连接...")
	//读取客户端发送的数据 阻塞
	buf:=make([]byte,4096)
	n,err:=conn.Read(buf)
	if err!=nil{
		fmt.Println("conn.Read err:",err)
		return
	}
	fmt.Println("服务器处理数据",string(buf[:n]))

	conn.Write(buf[:n])
	fmt.Println("服务器回发：",string(buf[:n]))

}
