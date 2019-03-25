package main

import (
	"github.com/astaxie/beego/logs"
	"logagent/util"
	"time"
)

func ServerRun() (err error) {
	for {
		msg := util.GetOneLine()
		err = SendToKafka(msg)
		if err != nil {
			logs.Error("send to kafka err", err)
			time.Sleep(time.Millisecond * 100)
			continue
		}
	}
	defer util.KafkaClient.Close()
	return
}

func SendToKafka(msg *util.TextMsg) (err error) {
	//logs.Debug("Msg:%s,Msg Topic:%s/n",msg.Msg,msg.Topic)
	//fmt.Printf("Msg:%s,Msg Topic:%s/n",msg.Msg,msg.Topic)
	util.SaveToKafka(msg)
	return
}
