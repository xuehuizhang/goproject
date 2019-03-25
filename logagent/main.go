package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
)
import "logagent/util"

func main() {
	fileName := "./conf/logagent.conf"
	err := util.LoadConf("ini", fileName) //加载配置文件
	if err != nil {
		fmt.Println("load conf err", err)
		panic("load conf err")
		return
	}

	err = util.InitLogger() //初始化日志
	if err != nil {
		fmt.Println("init logger err", err)
		panic("init logger err")
		return
	}

	err = util.InitTail(util.AppConfig.CollectConf, util.AppConfig.ChanSize) //初始化tail
	if err != nil {
		logs.Error("init tail err", err)
		return
	}

	err=util.InitKafka(util.AppConfig.Addr)
	if err != nil {
		logs.Error("init kafka err", err)
		return
	}

	go func() {
		for {
			logs.Debug("this is text err")
		}
	}()
	err = ServerRun()
	if err != nil {
		logs.Error("server run err")
		return
	}
}
