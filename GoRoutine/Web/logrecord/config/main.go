package main

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

func main(){
	conf,err:=config.NewConfig("ini","./Web/logrecord/config/logagent.conf")
	if err!=nil{
		fmt.Println("new config err",err)
		return
	}
	port,err:=conf.Int("server::listen_port")
	if err!=nil{
		fmt.Println("conf int err",err)
		return
	}
	fmt.Println("server port",port)

	level:=conf.String("logs::log_level")
	fmt.Println("logs level:",level)
}
