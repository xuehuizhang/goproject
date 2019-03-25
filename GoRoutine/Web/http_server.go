package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w,"hello")
	fmt.Println("hello")
}

func main()  {
	http.HandleFunc("/",hello)
	err:=http.ListenAndServe(":8080",nil)
	if err!=nil{
		fmt.Println("http listen failed")
	}
}
