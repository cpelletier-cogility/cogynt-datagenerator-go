package main

import (
	"bufio"
	"cogynt-datagenerator-go/datagenerator"
	"fmt"
	"os"
	"strings"
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
	fmt.Println("This is the start of the Data Generator Application.")

	for {
		reader := bufio.NewReader(os.Stdin)

		datagenerator.GenerateData()
		fmt.Println("Do you want to generate another data type? [Y/n]")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSuffix(userInput, "\n")
		if userInput == "Y" {
			continue
		} else {
			break
		}
	}
}
