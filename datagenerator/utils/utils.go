package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RequestItemCount(dataType string) int64 {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("How many %ss would you like to create?\n", dataType)
		countInput, _ := reader.ReadString('\n')
		countInput = strings.TrimSuffix(countInput, "\n")
		itemCount, err := strconv.ParseInt(countInput, 10, 64)
		if err == nil {
			return itemCount
		} else {
			fmt.Println("Please enter a valid integer")
		}
	}
}

func RequestTopicName(defaultTopic string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("What is the name of the topic you want to push to? (%s)\n", defaultTopic)
	topicInput, _ := reader.ReadString('\n')
	topicInput = strings.TrimSuffix(topicInput, "\n")
	if topicInput == "" {
		topicInput = defaultTopic
	}
	return topicInput
}

func RequestJsonFileName(defaultName string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("What is the name of the file you want to save to? (%s)\n", defaultName)
	topicInput, _ := reader.ReadString('\n')
	topicInput = strings.TrimSuffix(topicInput, "\n")
	topicInput = strings.TrimSuffix(topicInput, ".json")
	if topicInput == "" {
		topicInput = defaultName
	}
	return topicInput + ".json"
}
