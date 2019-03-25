package main

import (
	"FinalLogagent/model"
	"FinalLogagent/util"
)

func ServerRun() (err error) {
	for{
		msg:=util.GetOneLine()
		SendToKafka(msg)
	}
	defer util.KafkaClient.Close()
	return
}

func SendToKafka(msg *model.TextMsg)(err error){
	util.SaveToKafka(msg)
	return
}
