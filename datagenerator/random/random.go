package random

import (
	"github.com/brianvoe/gofakeit/v6"
)

func init() {
	faker := gofakeit.NewCrypto()
	gofakeit.SetGlobalFaker(faker)
}

type PointGeometry struct {
	Coordinates [2]float64 `json:"coordinates"`
	Type        string     `json:"type"`
}

type Feature struct {
	Geometry   PointGeometry `json:"geometry"`
	Type       string        `json:"type"`
	Properties struct{}      `json:"properties"`
}

type AddressInfo struct {
	Address  string  `json:"address"`
	Street   string  `json:"street"`
	City     string  `json:"city"`
	State    string  `json:"state"`
	Zip      string  `json:"zip"`
	Country  string  `json:"country"`
	Location Feature `json:"location"`
}

func GenerateAddress() AddressInfo {
	addressInfo := gofakeit.Address()
	return AddressInfo{
		Address: addressInfo.Address,
		Street:  addressInfo.Street,
		City:    addressInfo.City,
		State:   addressInfo.State,
		Zip:     addressInfo.Zip,
		Country: addressInfo.Country,
		Location: Feature{
			Geometry: PointGeometry{
				Coordinates: [2]float64{addressInfo.Longitude, addressInfo.Latitude},
				Type:        "Point",
			},
			Type:       "Feature",
			Properties: struct{}{},
		},
	}
}

type JobInfo struct {
	Id         string      `json:"id"`
	Descriptor string      `json:"descriptor"`
	Level      string      `json:"level"`
	Title      string      `json:"title"`
	Company    string      `json:"company"`
	Address    AddressInfo `json:"address"`
}

func GenerateRandomJob() JobInfo {
	jobInfo := gofakeit.Job()
	address := GenerateAddress()
	return JobInfo{
		Id:         gofakeit.UUID(),
		Descriptor: jobInfo.Descriptor,
		Level:      jobInfo.Level,
		Title:      jobInfo.Title,
		Company:    jobInfo.Company,
		Address:    address,
	}
}
