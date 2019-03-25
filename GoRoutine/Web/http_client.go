package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	res,err:=http.Get("https://www.baidu.com")
	if err!=nil{
		fmt.Println("http.Get err",err)
		return
	}
	buf,err:=ioutil.ReadAll(res.Body)
	if err!=nil{
		fmt.Println("ioutil.ReadAll err",err)
		return
	}
	fmt.Println(string(buf[:]))
}
