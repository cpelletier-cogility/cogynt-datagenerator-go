package kafkaproducer

import (
	"cogynt-datagenerator-go/datagenerator/config"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var kafkaConfig config.KafkaConfig

func init() {
	kafkaConfig = config.GetKafkaConfig()
	fmt.Printf("%+v", kafkaConfig)
}

func InitializeAsyncProducer() *kafka.Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaConfig.BootstrapServers,
		"client.id":         kafkaConfig.Producer.ClientId})
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Created Producer %v\n", p)
	return p
}

func HandleProducedMessages(producer *kafka.Producer) {
	for e := range producer.Events() {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
			} else {
				fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
			}
		}
	}
}
