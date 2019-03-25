package main

import (
	"fmt"
	"net"
	"os"
)

func main()  {
	conn,err:=net.Dial("tcp","172.18.2.34:8000")
	if err!=nil{
		fmt.Println("net.Dial error:",err)
		return
	}
	defer conn.Close()

	//获取用户的键盘输入(stdin)数据,发送给客户端
	go func() {
		buf:=make([]byte,4096)
		for{
			fmt.Print("Send:")
			n,err:=os.Stdin.Read(buf)
			if err!=nil{
				fmt.Println("os.Stdin.Read err",err)
				continue
			}
			conn.Write(buf[:n])
		}
	}()
	//回显服务器发的大写数据
	buf:=make([]byte,4096)
	for{
		n,err:=conn.Read(buf)
		if n==0{
			fmt.Println("检测到服务器退出,退出")
			return
		}
		if err!=nil{
			fmt.Println("conn.read err",err)
			continue
		}
		fmt.Println("读取服务端返回的大写数据：",string(buf[:n]))
	}
}
