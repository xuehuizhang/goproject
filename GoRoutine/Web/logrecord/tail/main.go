package main

import (
	"fmt"
	"github.com/hpcloud/tail"
)

func main()  {
	filename:="C:/Program Files (x86)/zookeeper-3.4.13/zookeeper-3.4.13.jar"
	tails,err:=tail.TailFile(filename,tail.Config{
		ReOpen:true,
		Follow:true,
		MustExist:false,
		Poll:true,
	})
	if err!=nil{
		fmt.Println("tail file err:",err)
		return
	}

	var msg *tail.Line
	var ok bool
	for true{
		msg,ok=<-tails.Lines
		if !ok{
			fmt.Println("file close")
			continue
		}
		fmt.Println("msg:",msg)
	}
}
