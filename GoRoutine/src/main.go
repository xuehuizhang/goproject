package main

import "fmt"

func main()  {
	var ch chan string
	ch=make(chan string)
	for i:=0;i<100;i++{
		go PrintStr(i,ch)
	}

	for{
		msg:=<-ch
		fmt.Println(msg)
	}
}

func PrintStr(i int,ch chan string)  {
	for{
		ch<-fmt.Sprintf("hello world from go routine %d",i)
	}
}
