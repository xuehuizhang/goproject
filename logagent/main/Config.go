package main

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/config"
)
var (
	appConfig *Config
)
type Config struct{
	LogLevel string
	LogPath string
	CollectConf []CollectConf
}
type CollectConf struct {
	LogPath string
	Topic string
}

func LoadConf(confType,fileName string)(err error){
	conf,err:=config.NewConfig(confType,fileName)
	if err!=nil{
		fmt.Println("config newconfig err",err)
		return
	}
	appConfig=&Config{}
	appConfig.LogPath=conf.String("logs::log_path")

	if len(appConfig.LogPath)==0{
		appConfig.LogPath="./logs"
	}

	appConfig.LogLevel=conf.String("logs::log_level")
	if len(appConfig.LogLevel)==0{
		appConfig.LogLevel="debug"
	}

	err=loadCollectConf(conf)
	return
}
func loadCollectConf(conf config.Configer)(err error) {
	cc:=CollectConf{}
	cc.LogPath=conf.String("collect::log_path")
	if len(cc.LogPath)==0{
		err=errors.New("invalid collec::logpath")
		return
	}
	cc.Topic=conf.String("collect::topic")
	if len(cc.Topic)==0{
		err=errors.New("invalid collec::topic")
		return
	}
	appConfig.CollectConf=append(appConfig.CollectConf,cc)
	return
}