package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type ProducerConfig struct {
	ClientId string `yaml:"client-id"`
}

type KafkaConfig struct {
	BootstrapServers string         `yaml:"bootstrap-servers"`
	Producer         ProducerConfig `yaml:"producer"`
}

type Config struct {
	Kafka KafkaConfig `yaml:"kafka"`
}

var config = Config{}

func init() {
	yamlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("Unable to read kafka producer config: #%v ", err)
	}
	err = yaml.Unmarshal([]byte(yamlFile), &config)
	if err != nil {
		log.Fatalf("Unable to parse config: %v", err)
	}
	fmt.Printf("%+v", config)
}

func GetConfig() Config {
	return config
}

func GetProducerConfig() ProducerConfig {
	return config.Kafka.Producer
}

func GetKafkaConfig() KafkaConfig {
	return config.Kafka
}
