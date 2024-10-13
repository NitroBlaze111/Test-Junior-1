package models

import "time"

// AirQualityReading represents a single air quality sensor reading
type AirQualityReading struct {
	SensorID  string    `json:"sensor_id"`
	Timestamp time.Time `json:"timestamp"`
	PM25      float64   `json:"pm25"`
	CO2       float64   `json:"co2"`
}
