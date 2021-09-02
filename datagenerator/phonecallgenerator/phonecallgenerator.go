package phonecallgenerator

import (
	"cogynt-datagenerator-go/datagenerator/random"
	"cogynt-datagenerator-go/datagenerator/utils"
	"math/rand"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/brianvoe/gofakeit/v6"
)

type PhoneCall struct {
	Id         string    `json:"id"`
	FromNumber string    `json:"from_number"`
	ToNumber   string    `json:"to_number"`
	Occurred   time.Time `json:"occurred"`
	Duration   int       `json:"duration"`
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
		utils.GenerateJsonData("phone_call", phoneCalls)
	case "kafka":
		utils.GenerateKafkaData("phone_call", phoneCalls)
	}
	return phoneCalls
}

func ShufflePeople(people []random.PersonInfo) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(people), func(i, j int) { people[i], people[j] = people[j], people[i] })
}

func PopPerson(people *[]random.PersonInfo) random.PersonInfo {
	f := len(*people)
	rv := (*people)[f-1]
	*people = append((*people)[:f-1])
	return rv
}

func GeneratePhoneCallDataWithPersons(outputType string, people []random.PersonInfo) []PhoneCall {
	phoneCallCount := utils.RequestItemCount("phone call")
	outside := false
	prompt := &survey.Confirm{
		Message: "Can calls randomly be made outside of the persons list?",
	}
	survey.AskOne(prompt, &outside)
	randomCaller := false
	prompt = &survey.Confirm{
		Message: "Do you want to have a little less random/crazy connections?",
	}
	survey.AskOne(prompt, &randomCaller)

	faker := gofakeit.New(0)
	var phoneCalls []PhoneCall
	var i int64
	shuffled_people := people
	ShufflePeople(shuffled_people)
	from := PopPerson(&shuffled_people).PhoneNumber
	for i = 0; i < phoneCallCount; i++ {
		var to string
		if outside && faker.Number(1, 2) == 1 {
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
		if randomCaller || len(people) <= 2 || faker.Number(1, 2) != 1 {
			shuffled_people = people
			ShufflePeople(shuffled_people)
			from = PopPerson(&shuffled_people).PhoneNumber
		} else {
			from = to
		}
	}
	switch outputType {
	case "json":
		utils.GenerateJsonData("phone_call", phoneCalls)
	case "kafka":
		utils.GenerateKafkaData("phone_call", phoneCalls)
	}
	return phoneCalls
}
