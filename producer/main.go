package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"os"
)

func main() {
	topicName := os.Getenv("TOPIC_NAME")
	fmt.Println("Topic name is " + topicName)
	producer, _ := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	msg := &sarama.ProducerMessage{
		Topic: topicName,
		Value: sarama.StringEncoder("hello kafka"),
	}

	partition, offset, _ := producer.SendMessage(msg)
	log.Printf("Sent to partion %v and the offset is %v", partition, offset)
}
