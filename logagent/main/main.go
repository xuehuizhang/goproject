package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
)
import "logagent/util"
func main()  {
	fileName:="./conf/logagent.conf"
	err:=util.LoadConf("ini",fileName)
	if err!=nil{
		fmt.Println("load conf err",err)
		panic("load conf err")
		return
	}

	err=util.InitLogger()
	if err!=nil{
		fmt.Println("init logger err",err)
		panic("init logger err")
		return
	}
	logs.Debug("this is a debug text")
}