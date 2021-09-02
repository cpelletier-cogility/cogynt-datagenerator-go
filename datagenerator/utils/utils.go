package utils

import (
	"bufio"
	"cogynt-datagenerator-go/datagenerator/kafkaproducer"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func RequestItemCount(dataType string) int64 {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("How many %ss would you like to create?\n", dataType)
		countInput, _ := reader.ReadString('\n')
		countInput = strings.TrimSuffix(countInput, "\n")
		itemCount, err := strconv.ParseInt(countInput, 10, 64)
		if err == nil {
			return itemCount
		} else {
			fmt.Println("Please enter a valid integer")
		}
	}
}

func RequestTopicName(defaultTopic string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("What is the name of the topic you want to push to? (%s)\n", defaultTopic)
	topicInput, _ := reader.ReadString('\n')
	topicInput = strings.TrimSuffix(topicInput, "\n")
	if topicInput == "" {
		topicInput = defaultTopic
	}
	return topicInput
}

func RequestJsonFileName(defaultName string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("What is the name of the file you want to save to? (%s)\n", defaultName)
	topicInput, _ := reader.ReadString('\n')
	topicInput = strings.TrimSuffix(topicInput, "\n")
	topicInput = strings.TrimSuffix(topicInput, ".json")
	if topicInput == "" {
		topicInput = defaultName
	}
	path := "./json_output"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0700)
		if err != nil {
			log.Fatalf("failed to create json_output directory: %s", err)
		}
	}
	return path + "/" + topicInput + ".json"
}

func GenerateJsonData(defaultName string, v interface{}) {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice:
		fileName := RequestJsonFileName(defaultName)
		file, err := os.OpenFile(fileName, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
		datawriter := bufio.NewWriter(file)
		if err != nil {
			log.Fatalf("failed creating file: %s", err)
		}
		fmt.Printf("Saving json to %v...\n", fileName)
		s := reflect.ValueOf(v)

		for i := 0; i < s.Len(); i++ {
			e := s.Index(i).Interface()
			entity, err := json.Marshal(e)
			if err != nil {
				fmt.Println(err)
				break
			}
			_, _ = datawriter.WriteString(string(entity) + "\n")
		}
		datawriter.Flush()
		file.Close()
	}
}

func GenerateKafkaData(defaultName string, v interface{}) {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice:
		topic := RequestTopicName(defaultName)
		fmt.Printf("Sending %v data to kafka...\n", topic)

		var p *kafka.Producer
		p = kafkaproducer.InitializeAsyncProducer()
		defer p.Close()

		go func() {
			for e := range p.Events() {
				switch ev := e.(type) {
				case *kafka.Message:
					if ev.TopicPartition.Error != nil {
						fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
					} else {
						fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
					}
				}
			}
		}()

		go kafkaproducer.HandleProducedMessages(p)

		s := reflect.ValueOf(v)

		for i := 0; i < s.Len(); i++ {
			e := s.Index(i).Interface()
			entity, err := json.Marshal(e)
			if err != nil {
				fmt.Println(err)
			} else {
				p.Produce(&kafka.Message{
					TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
					Value:          []byte(entity),
				}, nil)
			}
		}

		// Wait for message deliveries before shutting down
		p.Flush(15 * 1000)
	}
}
