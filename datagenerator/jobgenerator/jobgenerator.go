package jobgenerator

import (
	"bufio"
	"cogynt-datagenerator-go/datagenerator/kafkaproducer"
	"cogynt-datagenerator-go/datagenerator/random"
	"cogynt-datagenerator-go/datagenerator/utils"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func GenerateJsonData(jobs []random.JobInfo) {
	fileName := utils.RequestJsonFileName("job")
	file, err := os.OpenFile(fileName, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	datawriter := bufio.NewWriter(file)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	fmt.Println("Saving people to json file...")
	var i int
	for i = 0; i < len(jobs); i++ {
		person, err := json.Marshal(jobs[i])
		if err != nil {
			fmt.Println(err)
			break
		}
		_, _ = datawriter.WriteString(string(person) + "\n")
	}
	datawriter.Flush()
	file.Close()
}

func GenerateKafkaData(jobs []random.JobInfo) []random.JobInfo {
	topic := utils.RequestTopicName("job")
	fmt.Println("Sending job data to kafka...")

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

	for _, job := range jobs {
		job, err := json.Marshal(job)
		if err != nil {
			fmt.Println(err)
		} else {
			p.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Value:          []byte(job),
			}, nil)
		}
	}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
	return jobs
}

func GenerateJobData(outputType string) []random.JobInfo {
	jobCount := utils.RequestItemCount("job")
	var jobs []random.JobInfo

	var i int64
	for i = 0; i < jobCount; i++ {
		jobInfo := random.GenerateRandomJob()
		fmt.Printf("%+v\n", jobInfo)
		jobs = append(jobs, jobInfo)
	}
	switch outputType {
	case "json":
		GenerateJsonData(jobs)
	case "kafka":
		GenerateKafkaData(jobs)
	}
	return jobs
}
