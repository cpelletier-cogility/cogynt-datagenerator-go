package phonecallgenerator

import (
	"bufio"
	"cogynt-datagenerator-go/datagenerator/kafkaproducer"
	"cogynt-datagenerator-go/datagenerator/persongenerator"
	"cogynt-datagenerator-go/datagenerator/utils"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type PhoneCall struct {
	Id         string    `json:"id"`
	FromNumber string    `json:"from_number"`
	ToNumber   string    `json:"to_number"`
	Occurred   time.Time `json:"occurred"`
	Duration   int       `json:"duration"`
}

func GenrateJSONData(phoneCalls []PhoneCall) []PhoneCall {
	fileName := utils.RequestJsonFileName("phone_call")
	file, err := os.OpenFile(fileName, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	datawriter := bufio.NewWriter(file)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	fmt.Println("Saving phone calls to json file...")
	var i int
	for i = 0; i < len(phoneCalls); i++ {
		phoneCall, err := json.Marshal(phoneCalls[i])
		if err != nil {
			fmt.Println(err)
			break
		}
		_, _ = datawriter.WriteString(string(phoneCall) + "\n")
	}
	datawriter.Flush()
	file.Close()
	return phoneCalls
}

func GenerateKafkaData(phoneCalls []PhoneCall) []PhoneCall {
	topic := utils.RequestTopicName("phone_call")
	fmt.Println("Sending phone call data to kafka...")

	var p *kafka.Producer
	p = kafkaproducer.InitializeAsyncProducer()
	defer p.Close()

	go kafkaproducer.HandleProducedMessages(p)

	for _, call := range phoneCalls {
		call, err := json.Marshal(call)
		if err != nil {
			fmt.Println(err)
		} else {
			p.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Value:          []byte(call),
			}, nil)
		}
	}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
	return phoneCalls
}

func GeneratePhoneCallData(outputType string) []PhoneCall {
	phoneCallCount := utils.RequestItemCount("phone call")
	faker := gofakeit.New(0)
	var phoneCalls []PhoneCall
	var i int64
	for i = 0; i < phoneCallCount; i++ {
		occurred := faker.DateRange(time.Time{}, time.Now())
		c := PhoneCall{
			Id:         faker.UUID(),
			FromNumber: faker.PhoneFormatted(),
			ToNumber:   faker.PhoneFormatted(),
			Occurred:   occurred,
			Duration:   faker.Number(0, 300000),
		}
		phoneCalls = append(phoneCalls, c)
	}
	switch outputType {
	case "json":
		GenrateJSONData(phoneCalls)
	case "kafka":
		GenerateKafkaData(phoneCalls)
	}
	return phoneCalls
}

func ShufflePeople(people []persongenerator.PersonInfo) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(people), func(i, j int) { people[i], people[j] = people[j], people[i] })
}

func PopPerson(people *[]persongenerator.PersonInfo) persongenerator.PersonInfo {
	f := len(*people)
	rv := (*people)[f-1]
	*people = append((*people)[:f-1])
	return rv
}

func GeneratePhoneCallDataWithPersons(outputType string, people []persongenerator.PersonInfo) []PhoneCall {
	reader := bufio.NewReader(os.Stdin)
	phoneCallCount := utils.RequestItemCount("phone call")
	fmt.Println("Can calls randomly be made outside of the persons list? [Y/n]")
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSuffix(userInput, "\n")
	fmt.Println("Do you want to have a little less random/crazy connections? [Y/n]")
	randomCallerInput, _ := reader.ReadString('\n')
	randomCallerInput = strings.TrimSuffix(userInput, "\n")
	faker := gofakeit.New(0)
	var phoneCalls []PhoneCall
	var i int64
	shuffled_people := people
	ShufflePeople(shuffled_people)
	from := PopPerson(&shuffled_people).PhoneNumber
	for i = 0; i < phoneCallCount; i++ {
		var to string
		if userInput == "Y" && faker.Number(1, 2) == 1 {
			to = faker.PhoneFormatted()
		} else {
			to = PopPerson(&shuffled_people).PhoneNumber
		}
		occurred := faker.DateRange(time.Time{}, time.Now())
		c := PhoneCall{
			Id:         faker.UUID(),
			FromNumber: from,
			ToNumber:   to,
			Occurred:   occurred,
			Duration:   faker.Number(0, 300000),
		}
		phoneCalls = append(phoneCalls, c)
		if randomCallerInput != "Y" || len(people) <= 2 || faker.Number(1, 2) != 1 {
			shuffled_people = people
			ShufflePeople(shuffled_people)
			from = PopPerson(&shuffled_people).PhoneNumber
		} else {
			from = to
		}
	}
	switch outputType {
	case "json":
		GenrateJSONData(phoneCalls)
	case "kafka":
		GenerateKafkaData(phoneCalls)
	}
	return phoneCalls
}
