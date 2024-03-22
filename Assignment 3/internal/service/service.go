package service

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"

	"assignment-3/internal/model"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GenerateRandomStatus() model.Status {
	rand.Seed(time.Now().UnixNano())
	water := rand.Intn(100) + 1
	wind := rand.Intn(100) + 1
	return model.Status{
		Water: water,
		Wind:  wind,
	}
}

func (s *Service) DetermineStatus(status model.Status) string {
	waterStatus := ""
	windStatus := ""

	if status.Water < 5 {
		waterStatus = "Aman"
	} else if status.Water >= 6 && status.Water <= 8 {
		waterStatus = "Siaga"
	} else {
		waterStatus = "Bahaya"
	}

	if status.Wind < 6 {
		windStatus = "Aman"
	} else if status.Wind >= 7 && status.Wind <= 15 {
		windStatus = "Siaga"
	} else {
		windStatus = "Bahaya"
	}

	return "Status Air: " + waterStatus + ", Status Angin: " + windStatus
}

func (s *Service) UpdateJSONStatus(status model.Status, message string) error {
	data := map[string]interface{}{
		"status":  status,
		"message": message,
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("status.json", jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
