package persistence

import (
	"encoding/csv"
	"os"
	"strconv"
)

// SaveAverageToCSV saves the average pollutant values to a CSV file
func SaveAverageToCSV(averages map[string]float64, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write headers
	writer.Write([]string{"Pollutant", "Average Value"})

	// Write data
	for pollutant, avgValue := range averages {
		writer.Write([]string{pollutant, strconv.FormatFloat(avgValue, 'f', 2, 64)})
	}

	return nil
}

// SaveToCSV saves the highest pollutant data by hour to a CSV file
func SaveToCSV(data map[int]string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write headers
	writer.Write([]string{"Hour", "Highest Pollutant"})

	// Write data
	for hour, pollutant := range data {
		writer.Write([]string{strconv.Itoa(hour), pollutant})
	}

	return nil
}
