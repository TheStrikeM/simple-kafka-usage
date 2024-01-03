package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"os"
)

func main() {
	topicName := os.Getenv("TOPIC_NAME")
	fmt.Println("Topic name is " + topicName)

	admin, _ := sarama.NewClusterAdmin([]string{"localhost:9092"}, nil)
	defer admin.Close()

	admin.CreateTopic(topicName, &sarama.TopicDetail{
		NumPartitions:     1,
		ReplicationFactor: 1,
	}, false)

}
