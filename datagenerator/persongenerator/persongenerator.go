package persongenerator

import (
	"bufio"
	"cogynt-datagenerator-go/datagenerator/kafkaproducer"
	"cogynt-datagenerator-go/datagenerator/utils"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type PointGeometry struct {
	Coordinates [2]float64 `json:"coordinates"`
	Type        string     `json:"type"`
}

type Feature struct {
	Geometry   PointGeometry `json:"geometry"`
	Type       string        `json:"type"`
	Properties struct{}      `json:"properties"`
}

type PersonInfo struct {
	Name        string  `json:"name"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	Gender      string  `json:"gender"`
	Email       string  `json:"email"`
	Country     string  `json:"country"`
	City        string  `json:"city"`
	PostalCode  string  `json:"postalcode"`
	State       string  `json:"state"`
	Longitude   float64 `json:"lon"`
	Latitude    float64 `json:"lat"`
	Location    Feature `json:"loc"`
	Id          string  `json:"id"`
	PhoneNumber string  `json:"phone_number"`
}

func GenrateJSONData(people []PersonInfo) {
	fileName := utils.RequestJsonFileName("person")
	file, err := os.OpenFile(fileName, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	datawriter := bufio.NewWriter(file)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	fmt.Println("Saving people to json file...")
	var i int
	for i = 0; i < len(people); i++ {
		person, err := json.Marshal(people[i])
		if err != nil {
			fmt.Println(err)
			break
		}
		_, _ = datawriter.WriteString(string(person) + "\n")
	}
	datawriter.Flush()
	file.Close()
}

func GenerateKafkaData(people []PersonInfo) []PersonInfo {
	topic := utils.RequestTopicName("person")
	fmt.Println("Sending person data to kafka...")

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

	for _, person := range people {
		person, err := json.Marshal(person)
		if err != nil {
			fmt.Println(err)
		} else {
			p.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Value:          []byte(person),
			}, nil)
		}
	}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
	return people
}

func GeneratePersonData(outputType string) []PersonInfo {
	personCount := utils.RequestItemCount("person")
	faker := gofakeit.New(0)

	var people []PersonInfo

	var i int64
	for i = 0; i < personCount; i++ {
		firstName := faker.FirstName()
		lastName := faker.LastName()
		address := faker.Address()

		p := PersonInfo{
			Name:       firstName + " " + lastName,
			FirstName:  firstName,
			LastName:   lastName,
			Gender:     faker.Gender(),
			Email:      faker.Email(),
			Country:    address.Country,
			City:       address.City,
			PostalCode: address.Zip,
			State:      address.State,
			Longitude:  address.Longitude,
			Latitude:   address.Latitude,
			Location: Feature{
				Geometry: PointGeometry{
					Coordinates: [2]float64{address.Longitude, address.Latitude},
					Type:        "Point",
				},
				Type:       "Feature",
				Properties: struct{}{},
			},
			Id:          faker.UUID(),
			PhoneNumber: faker.PhoneFormatted(),
		}
		people = append(people, p)
	}
	switch outputType {
	case "json":
		GenrateJSONData(people)
	case "kafka":
		GenerateKafkaData(people)
	}
	return people
}
