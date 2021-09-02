package jobgenerator

import (
	"cogynt-datagenerator-go/datagenerator/random"
	"cogynt-datagenerator-go/datagenerator/utils"
)

func GenerateJobData(outputType string) []random.JobInfo {
	jobCount := utils.RequestItemCount("job")
	var jobs []random.JobInfo

	var i int64
	for i = 0; i < jobCount; i++ {
		jobInfo := random.GenerateRandomJob()
		jobs = append(jobs, jobInfo)
	}
	switch outputType {
	case "json":
		utils.GenerateJsonData("job", jobs)
	case "kafka":
		utils.GenerateKafkaData("job", jobs)
	}
	return jobs
}
