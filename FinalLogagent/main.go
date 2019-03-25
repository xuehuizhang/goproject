package main

import (
	"FinalLogagent/util"
	"github.com/astaxie/beego/logs"
)

func main()  {
	fileName:="./conf/logagent.conf"
	err:=util.LoadConf("ini",fileName)  //首次执行 加载配置文件
	if err!=nil{
		panic("load conf err")
		return
	}

	err=util.InitLogger()                        //初始化日志组件
	if err!=nil{
		logs.Error("init logger err",err)
		return
	}

	//初始化tail组件
	err=util.InitTail(util.AppConfig.CollectConf,util.AppConfig.ChanSize)
	if err!=nil{
		logs.Error("init tail err",err)
		return
	}

	err=util.InitKafka(util.AppConfig.KafkaAddr,util.AppConfig.KafkaPort)
	if err!=nil{
		logs.Error("init kafka err",err)
		return
	}

	go func() {
		for{
			logs.Debug("this is text err")
		}
	}()

	err=ServerRun()
	if err!=nil{
		logs.Error("server run err",err)
		return
	}

}
