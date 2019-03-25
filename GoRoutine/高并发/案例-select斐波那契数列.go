package main

import "fmt"

func main()  {
	ch:=make(chan int)
	quit:=make(chan bool)

	go fibn(ch,quit,20)

	for{
		select {
		case num:=<-ch:
			fmt.Print(num," ")
			case <-quit:
				return
		}
	}
}

func fibn(ch chan<- int,quit chan <- bool,n int){
	x,y:=1,1
	for i:=1;i<20;i++{
		ch<-x
		x,y=y,x+y
	}
	quit<-true
}