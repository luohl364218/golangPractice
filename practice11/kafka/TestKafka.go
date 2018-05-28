package main

import (
	"github.com/Shopify/sarama"
	"fmt"
)

func main() {
	config := sarama.NewConfig()
	//生产者配置
	//需要ACK
	config.Producer.RequiredAcks = sarama.WaitForAll
	//做分区 随机
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}

	defer client.Close()

	msg := &sarama.ProducerMessage{}
	msg.Topic = "nginx_log"
	msg.Value = sarama.StringEncoder("this is a test, I am good")

	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}

	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}


