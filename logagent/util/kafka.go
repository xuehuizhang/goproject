package util

import (
	"fmt"
	"github.com/shopify/sarama"
)
var (
	KafkaClient sarama.SyncProducer
)
func InitKafka(addr string) (err error){
	config:=sarama.NewConfig()
	config.Producer.RequiredAcks=sarama.WaitForAll     //ack
	config.Producer.Partitioner=sarama.NewRandomPartitioner  //随机分区
	config.Producer.Return.Successes=true
	client,syncerr:=sarama.NewSyncProducer([]string{addr},config) //"172.18.2.34:9092"
	if err!=nil{
		fmt.Println("producer close,err:",err)
		err=syncerr
		return
	}
	KafkaClient=client
	return
}

func SaveToKafka(textMsg *TextMsg) (err error) {


	msg:=&sarama.ProducerMessage{}
	msg.Topic=textMsg.Topic
	msg.Value=sarama.StringEncoder(textMsg.Msg)

		pid, offset, kafkaErr := KafkaClient.SendMessage(msg)
		if err != nil {
			fmt.Println("send mssage failed", err)
			err=kafkaErr
			return
		}
		fmt.Printf("pid:%v offset:%v\n", pid, offset)
   return
}