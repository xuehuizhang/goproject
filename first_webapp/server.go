package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter,request *http.Request){
	fmt.Fprintf(writer,"hello world %s",request.URL.Path[1:])
}

func main()  {
	/*http.HandleFunc("/",handler)   //
	http.ListenAndServe(":8080",nil)*/ //这里监听8080端口，并且使用默认的多路复用器

	//1 创建多路复用器
	mus:=http.NewServeMux()

	//处理静态文件
	/*files:=http.FileServer(http.Dir("/public"))
	mus.Handle("/static/",http.StripPrefix("/static/",files))*/
	mus.HandleFunc("/",handler)

	server:=&http.Server{
		Addr:":8080",
		Handler:mus,
	}
	server.ListenAndServe()
}
