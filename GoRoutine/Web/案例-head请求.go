package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

var url=[]string{
	"http://www.baidu.com",
	"http://google.com",
	"http://taobao.com",
}

func main()  {
	for _,v:=range url{
		c:=http.Client{
			Transport:&http.Transport{
				Dial: func(network, addr string) (conn net.Conn, e error) {
					timeOut:=time.Microsecond*2  //自己指定超时时间
					return net.DialTimeout(network,addr,timeOut)
				},
			},
		}
        //用自己生成一个客户端请求
		res,err:=c.Head(v)
		if err!=nil{
			fmt.Println("http.Head err",err)
			continue
		}
		fmt.Println("head succ status=",res.Status)
	}
}
