package jobgenerator

type JobInfo struct {
	Id         string `json:"id"`
	Descriptor string `json:"descriptor"`
	Level      string `json:"level"`
	Title      string `json:"title"`
}

func GenerateJobData(outputType string) []JobInfo {
	var jobs []JobInfo
	return jobs
}
