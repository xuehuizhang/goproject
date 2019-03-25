package main

import (
	"fmt"
	"io"
	"net/http"
)

func main()  {
	fs,err:=http.Get("https://www.baidu.com")
	if err!=nil{
		fmt.Println("Http Get Error",err)
		return
	}

	defer fs.Body.Close()

	buf :=make([]byte,4096)
	var result string
	for{
		n,err:=fs.Body.Read(buf)
		if n==0{
			break
		}
		if err!=nil&&err!=io.EOF{
			fmt.Println(err)
			break
		}
		result+=string(buf[:n])
	}
	fmt.Println(result)
}
