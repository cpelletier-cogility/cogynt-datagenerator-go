package main

import (
	"cogynt-datagenerator-go/datagenerator"

	"github.com/AlecAivazis/survey/v2"
)

func inputInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func main() {
	for {
		datagenerator.GenerateData()
		createAnother := false
		prompt := &survey.Confirm{
			Message: "Do you want to generate another data type?",
		}
		survey.AskOne(prompt, &createAnother)
		if createAnother {
			continue
		} else {
			break
		}
	}
}
