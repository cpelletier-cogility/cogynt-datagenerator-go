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

func GenerateRandomAddress() AddressInfo {
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
	address := GenerateRandomAddress()
	return JobInfo{
		Id:         gofakeit.UUID(),
		Descriptor: jobInfo.Descriptor,
		Level:      jobInfo.Level,
		Title:      jobInfo.Title,
		Company:    jobInfo.Company,
		Address:    address,
	}
}

type PersonInfo struct {
	Name        string  `json:"name"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	Gender      string  `json:"gender"`
	Email       string  `json:"email"`
	Country     string  `json:"country"`
	City        string  `json:"city"`
	PostalCode  string  `json:"postalcode"`
	State       string  `json:"state"`
	Longitude   float64 `json:"lon"`
	Latitude    float64 `json:"lat"`
	Location    Feature `json:"loc"`
	Id          string  `json:"id"`
	PhoneNumber string  `json:"phone_number"`
	JobId       string  `json:"job_id"`
}

func GenerateRandomPerson() PersonInfo {
	firstName := gofakeit.FirstName()
	lastName := gofakeit.LastName()
	address := GenerateRandomAddress()
	location := address.Location
	latitude := location.Geometry.Coordinates[1]
	longitude := location.Geometry.Coordinates[0]

	return PersonInfo{
		Name:        firstName + " " + lastName,
		FirstName:   firstName,
		LastName:    lastName,
		Gender:      gofakeit.Gender(),
		Email:       gofakeit.Email(),
		Country:     address.Country,
		City:        address.City,
		PostalCode:  address.Zip,
		State:       address.State,
		Longitude:   longitude,
		Latitude:    latitude,
		Location:    location,
		Id:          gofakeit.UUID(),
		PhoneNumber: gofakeit.PhoneFormatted(),
	}
}
