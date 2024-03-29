package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func main()  {
	config:=make(map[string]interface{})
	config["fileName"]="./web/logrecord/logcollect.log"
	config["level"]=logs.LevelDebug

	configStr,err:=json.Marshal(config)
	if err!=nil{
		fmt.Println("marshal failed,err",err)
		return
	}

	logs.SetLogger(logs.AdapterFile,string(configStr))

	logs.Debug("this is debug test,my name is %s","stu01")
}
