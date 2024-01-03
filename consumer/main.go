package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"os"
	"sync"
)

func main() {
	topicName := os.Getenv("TOPIC_NAME")
	fmt.Println("Topic name is " + topicName)

	consumer, _ := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	defer consumer.Close()

	partitionList, _ := consumer.Partitions(topicName)
	var wg sync.WaitGroup

	for _, partition := range partitionList {
		pc, _ := consumer.ConsumePartition("test", partition, sarama.OffsetOldest)
		wg.Add(1)
		go func(pc sarama.PartitionConsumer) {
			defer wg.Done()
			for message := range pc.Messages() {
				log.Printf("received message %v\n", string(message.Value))
			}
		}(pc)
	}

	wg.Wait()
}
