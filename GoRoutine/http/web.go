package main

import (
	"fmt"
	"net/http"
	"os"
)

func OpenSendFile(url string,w http.ResponseWriter)  {
	fs,err:=os.Open(url)
	if err!=nil{
		fmt.Println("os open file err",err)
		w.Write([]byte("no such file or directory"))
		return
	}
	buf:=make([]byte,4096)
	for{
		n,_:=fs.Read(buf)
		if n==0{
			return
		}
		w.Write(buf[:n])
	}
}

func handler(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("客户端请求：",r.URL)
	url:="C:/Users/miqil/Desktop/"+r.URL.String()
	OpenSendFile(url,w)
}

//web服务器练习
func main()  {
	//注册回调函数
	http.HandleFunc("/",handler)

	http.ListenAndServe(":8080",nil)
}
