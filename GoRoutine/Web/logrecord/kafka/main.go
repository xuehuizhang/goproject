package main

import (
	"fmt"
	"github.com/shopify/sarama"
)

func main(){
	fmt.Println("kafka")
	config:=sarama.NewConfig()
	config.Producer.RequiredAcks=sarama.WaitForAll     //ack
	config.Producer.Partitioner=sarama.NewRandomPartitioner  //随机分区
	config.Producer.Return.Successes=true
	client,err:=sarama.NewSyncProducer([]string{"172.18.2.34:9092"},config)
	if err!=nil{
		fmt.Println("producer close,err:",err)
		return
	}
	defer client.Close()

	msg:=&sarama.ProducerMessage{}
	msg.Topic="test"
	msg.Value=sarama.StringEncoder("this is a good test,my message is good")
    for {
		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("send mssage failed", err)
			return
		}
		fmt.Printf("pid:%v offset:%v\n", pid, offset)
	}
}
