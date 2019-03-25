package util

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
)
var (
	AppConfig *Config
)
type Config struct{
	LogLevel string
	LogPath string
	CollectConf []CollectConf

	ChanSize int
	Addr string
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
	AppConfig=&Config{}
	AppConfig.LogPath=conf.String("logs::log_path")

	if len(AppConfig.LogPath)==0{
		AppConfig.LogPath="./logs"
	}

	AppConfig.LogLevel=conf.String("logs::log_level")
	if len(AppConfig.LogLevel)==0{
		AppConfig.LogLevel="debug"
	}

	AppConfig.Addr=conf.String("kafka::addr")
	if len(AppConfig.Addr)==0{
		logs.Debug("kafka addr err")
		return
	}

	chanSize,err:=conf.Int("logs::chan_size")
	if err!=nil{
		AppConfig.ChanSize=100
	}
	AppConfig.ChanSize=chanSize

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
	AppConfig.CollectConf=append(AppConfig.CollectConf,cc)
	return
}