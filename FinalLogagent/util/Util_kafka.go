package util

import (
	"FinalLogagent/model"
	"fmt"
	"github.com/shopify/sarama"
)

var (
	KafkaClient sarama.SyncProducer
)

func InitKafka(addr string,port int)(err error)  {
	addr=fmt.Sprintf("%v:%v",addr,port)
	config:=sarama.NewConfig()
	config.Producer.RequiredAcks=sarama.WaitForAll
	config.Producer.Partitioner=sarama.NewRandomPartitioner
	config.Producer.Return.Successes=true
	client,syncerr:=sarama.NewSyncProducer([]string{addr},config)
	if syncerr!=nil{
		err=syncerr
		return
	}
	KafkaClient=client
	return
}

func SaveToKafka(textMsg *model.TextMsg)(err error){
	msg:=&sarama.ProducerMessage{}
	msg.Topic=textMsg.Topic
	msg.Value=sarama.StringEncoder(textMsg.Msg)

	pid, offset, kafkaErr := KafkaClient.SendMessage(msg)
	if err != nil {
		err=kafkaErr
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
	return
}
