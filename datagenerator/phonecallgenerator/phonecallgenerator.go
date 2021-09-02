package phonecallgenerator

import (
	"cogynt-datagenerator-go/datagenerator/random"
	"cogynt-datagenerator-go/datagenerator/utils"
	"math/rand"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/brianvoe/gofakeit/v6"
)

func GeneratePhoneCallData(outputType string) []random.PhoneCallInfo {
	phoneCallCount := utils.RequestItemCount("phone call")
	var phoneCalls []random.PhoneCallInfo
	var i int64
	for i = 0; i < phoneCallCount; i++ {
		phoneCalls = append(phoneCalls, random.GenerateRandomPhoneCall())
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

func GeneratePhoneCallDataWithPersons(outputType string, people []random.PersonInfo) []random.PhoneCallInfo {
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
	var phoneCalls []random.PhoneCallInfo
	var i int64
	shuffled_people := people
	ShufflePeople(shuffled_people)
	from := PopPerson(&shuffled_people).PhoneNumber
	for i = 0; i < phoneCallCount; i++ {
		var to string
		phoneCall := random.GenerateRandomPhoneCall()
		if !outside && faker.Number(1, 2) != 1 {
			phoneCall.ToNumber = PopPerson(&shuffled_people).PhoneNumber
		}
		phoneCall.FromNumber = from
		phoneCalls = append(phoneCalls, phoneCall)
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
