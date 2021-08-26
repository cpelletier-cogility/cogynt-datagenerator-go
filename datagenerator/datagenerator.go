package datagenerator

import (
	"cogynt-datagenerator-go/datagenerator/jobgenerator"
	"cogynt-datagenerator-go/datagenerator/persongenerator"
	"cogynt-datagenerator-go/datagenerator/phonecallgenerator"
	"cogynt-datagenerator-go/datagenerator/random"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

var dataTypes []string = []string{"jobs", "person", "phone_call"}

type DataMap struct {
	Person    []persongenerator.PersonInfo
	PhoneCall []phonecallgenerator.PhoneCall
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
		fmt.Printf("%+v\n", dataType)
		switch dataType {
		case "jobs":
			dataMap.Job = jobgenerator.GenerateJobData(outputType)
		case "person":
			// if dataMap.Job != nil {
			// 	employPeople := ""
			// 	prompt := &survey.Select{
			// 			Message: "Choose a color:",
			// 			Options: []string{"red", "blue", "green"},
			// 	}
			// 	survey.AskOne(prompt, &employPeople)
			// }
			people := persongenerator.GeneratePersonData(outputType)
			dataMap.Person = people
		case "phone_call":
			var phoneCalls []phonecallgenerator.PhoneCall
			if dataMap.Person != nil {
				phoneCalls = phonecallgenerator.GeneratePhoneCallDataWithPersons(outputType, dataMap.Person)
			} else {
				phoneCalls = phonecallgenerator.GeneratePhoneCallData(outputType)
			}
			dataMap.PhoneCall = phoneCalls
		}
	}
}
