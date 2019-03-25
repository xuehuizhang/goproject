package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func Hello(w http.ResponseWriter,r *http.Request)  {
	i:=1
	v:=0
	c:=i/v
	io.WriteString(w,"hello world"+strconv.Itoa(c))
}

func main()  {
	http.HandleFunc("/",logPanics(Hello))
	http.ListenAndServe(":8080",nil)
}

func logPanics(handle http.HandlerFunc)http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if x:=recover();x!=nil{
				log.Printf("[%v] caught panic:%v",request.RemoteAddr,x)
			}
		}()
		handle(writer,request)
	}
}