package main

import (
	"fmt"
	"net"
)

func main()  {
	conn,err:=net.Dial("tcp","127.0.0.1:8000")
	if err!=nil{
		fmt.Println("net.Dail err:",err)
		return
	}
	defer conn.Close()
    for {
    	var msg string
    	fmt.Println("Send:")
    	fmt.Scan(&msg)
		conn.Write([]byte(msg))

		buf := make([]byte, 4096)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("服务器回发错误")
			return
		}
		fmt.Println("服务器回发数据：", string(buf[:n]))
	}
}
