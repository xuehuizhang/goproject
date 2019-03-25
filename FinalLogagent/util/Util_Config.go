package util

import (
	"FinalLogagent/model"
	"fmt"
	"github.com/astaxie/beego/config"
)
var (
	AppConfig *model.Config
)
func LoadConf(confType,fileName string ) (err error){
	fmt.Println(confType,fileName)
	conf,err:=config.NewConfig(confType,fileName)
	if err!=nil{
		return
	}
	AppConfig=&model.Config{}
	AppConfig.LogPath=conf.String("logs::log_path")
	if len(AppConfig.LogPath)==0{
		AppConfig.LogPath="./logs"
	}

	AppConfig.LogLevel=conf.String("logs::log_level")
	if len(AppConfig.LogLevel)==0{
		AppConfig.LogLevel="debug"
	}

	AppConfig.KafkaAddr=conf.String("kafka::kafka_addr")
	if len(AppConfig.KafkaAddr)==0{
		err=fmt.Errorf("kafka addr nil")
		return
	}

	AppConfig.KafkaPort,err=conf.Int("kafka::kafka_port")
	if err!=nil{
		return
	}

	AppConfig.ChanSize,err=conf.Int("common::chan_size")
	if err!=nil{
		AppConfig.ChanSize=100
	}
	loadCollectConf(conf)
	return
}

func loadCollectConf(conf config.Configer){
	cc:=model.CollectConf{}
	cc.LogPath=conf.String("collect::log_path")
	if len(cc.LogPath)==0{
		cc.LogPath="./logs"
	}
	cc.Topic=conf.String("collect::topic")
	if len(cc.Topic)==0{
		cc.Topic="test"
	}
	AppConfig.CollectConf=append(AppConfig.CollectConf,cc)
	return
}