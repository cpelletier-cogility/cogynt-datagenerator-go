package datagenerator

import (
	"cogynt-datagenerator-go/datagenerator/jobgenerator"
	"cogynt-datagenerator-go/datagenerator/persongenerator"
	"cogynt-datagenerator-go/datagenerator/phonecallgenerator"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

var dataTypes []string = []string{"jobs", "person", "phone_call"}

type DataMap struct {
	Person    []persongenerator.PersonInfo
	PhoneCall []phonecallgenerator.PhoneCall
	Job       []jobgenerator.JobInfo
}

func GenerateData() {
	dataTypesSelections := []string{}
	dataTypePrompt := &survey.MultiSelect{
		Message: "Which event types would you like to create data for?",
		Options: dataTypes,
		Default: []string{"person"},
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
			fmt.Println("Jobs data creation has not been implemented yet")
		case "person":
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
