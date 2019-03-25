package util

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
)

func ConvertLogLevel(level string)int{
	switch level {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	}
	return logs.LevelDebug
}

func InitLogger()(err error) {
	config:=make(map[string]interface{})
	config["fileName"]=AppConfig.LogPath
	config["level"]=ConvertLogLevel(AppConfig.LogLevel)

	configStrings,err:=json.Marshal(config)
	if err!=nil{
		return
	}
	logs.SetLogger(logs.AdapterFile,string(configStrings))
	return
}
