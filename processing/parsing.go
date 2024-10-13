package processing

import (
	"air-quality-analyzer/models"
	"encoding/json"
	"errors"
)

// ParseReadings takes JSON data and parses it into AirQualityReading structs
func ParseReadings(data []byte) ([]models.AirQualityReading, error) {
	var readings []models.AirQualityReading
	err := json.Unmarshal(data, &readings)
	if err != nil {
		return nil, errors.New("error unmarshalling JSON data")
	}
	return readings, nil
}
