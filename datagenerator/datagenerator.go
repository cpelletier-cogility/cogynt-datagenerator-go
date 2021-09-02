package datagenerator

import (
	"cogynt-datagenerator-go/datagenerator/jobgenerator"
	"cogynt-datagenerator-go/datagenerator/persongenerator"
	"cogynt-datagenerator-go/datagenerator/phonecallgenerator"
	"cogynt-datagenerator-go/datagenerator/random"

	"github.com/AlecAivazis/survey/v2"
)

var dataTypes []string = []string{"jobs", "person", "phone_call"}

type DataMap struct {
	Person    []random.PersonInfo
	PhoneCall []random.PhoneCallInfo
	Job       []random.JobInfo
}

func GenerateData() {
	dataTypesSelections := []string{}
	dataTypePrompt := &survey.MultiSelect{
		Message: "Which event types would you like to create data for?",
		Options: dataTypes,
		Default: []string{},
	}
	survey.AskOne(dataTypePrompt, &dataTypesSelections)

	outputType := "json"
	outputTypePrompt := &survey.Select{
		Message: "What type of data do you want to produce?",
		Options: []string{"json", "kafka"},
	}
	survey.AskOne(outputTypePrompt, &outputType, survey.WithValidator(survey.Required))
	var dataMap DataMap
	for i := 0; i < len(dataTypesSelections); i++ {
		dataType := dataTypesSelections[i]
		switch dataType {
		case "jobs":
			dataMap.Job = jobgenerator.GenerateJobData(outputType)
		case "person":
			var people []random.PersonInfo
			if dataMap.Job != nil {
				employPeople := false
				prompt := &survey.Confirm{
					Message: "Should the persons be employed using the person topic?",
				}
				survey.AskOne(prompt, &employPeople)
				if employPeople {
					people = persongenerator.GeneratePersonWithJobData(outputType, dataMap.Job)
				} else {
					people = persongenerator.GeneratePersonData(outputType)
				}
			} else {
				people = persongenerator.GeneratePersonData(outputType)
			}
			dataMap.Person = people
		case "phone_call":
			var phoneCalls []random.PhoneCallInfo
			if dataMap.Person != nil {
				phoneCalls = phonecallgenerator.GeneratePhoneCallDataWithPersons(outputType, dataMap.Person)
			} else {
				phoneCalls = phonecallgenerator.GeneratePhoneCallData(outputType)
			}
			dataMap.PhoneCall = phoneCalls
		}
	}
}
