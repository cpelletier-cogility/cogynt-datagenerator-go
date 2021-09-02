package persongenerator

import (
	"cogynt-datagenerator-go/datagenerator/random"
	"cogynt-datagenerator-go/datagenerator/utils"

	"github.com/brianvoe/gofakeit/v6"
)

func GeneratePersonData(outputType string) []random.PersonInfo {
	personCount := utils.RequestItemCount("person")

	var people []random.PersonInfo

	var i int64
	for i = 0; i < personCount; i++ {
		people = append(people, random.GenerateRandomPerson())
	}
	switch outputType {
	case "json":
		utils.GenerateJsonData("person", people)
	case "kafka":
		utils.GenerateKafkaData("person", people)
	}
	return people
}

func GeneratePersonWithJobData(outputType string, jobs []random.JobInfo) []random.PersonInfo {
	personCount := utils.RequestItemCount("person")

	var people []random.PersonInfo

	var i int64
	for i = 0; i < personCount; i++ {
		person := random.GenerateRandomPerson()
		shuffled_jobs := jobs
		gofakeit.ShuffleAnySlice(shuffled_jobs)
		index := gofakeit.Number(0, len(shuffled_jobs)-1)
		person.JobId = shuffled_jobs[index].Id
		people = append(people, person)
	}
	switch outputType {
	case "json":
		utils.GenerateJsonData("person", people)
	case "kafka":
		utils.GenerateKafkaData("person", people)
	}
	return people
}
