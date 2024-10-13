package main

import (
	"air-quality-analyzer/analysis"
	"air-quality-analyzer/persistence"
	"air-quality-analyzer/processing"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// โหลดข้อมูลจาก mockdata.json
	data, err := ioutil.ReadFile("data/mockdata.json")
	if err != nil {
		log.Fatalf("Error reading mockdata.json: %v", err)
	}

	// Parse JSON เป็น AirQualityReading structs
	readings, err := processing.ParseReadings(data)
	if err != nil {
		log.Fatalf("Error parsing data: %v", err)
	}

	// คำนวณค่าเฉลี่ยของแต่ละสารมลพิษ
	averages := analysis.CalculateAverage(readings)
	fmt.Println("Average values for all pollutants:", averages)

	// หาสารมลพิษที่มีค่าเฉลี่ยสูงสุดในแต่ละชั่วโมง
	highestPollutantByHour, averagePollutants := analysis.FindHighestPollutantByHour(readings)
	fmt.Println("Highest Pollutant By Hour with values:")
	for hour, pollutant := range highestPollutantByHour {
		fmt.Printf("Hour: %d, Pollutant: %s, Average: %.2f\n", hour, pollutant, averagePollutants[hour][pollutant])
	}

	// บันทึกค่าเฉลี่ยลงใน CSV
	err = persistence.SaveAverageToCSV(averages, "average_values.csv")
	if err != nil {
		log.Fatalf("Error saving average values to CSV: %v", err)
	} else {
		fmt.Println("Average values saved to average_values.csv")
	}

	// บันทึกสารมลพิษที่มีค่าสูงสุดในแต่ละชั่วโมงลงใน CSV
	err = persistence.SaveToCSV(highestPollutantByHour, "highest_pollutants_by_hour.csv")
	if err != nil {
		log.Fatalf("Error saving highest pollutants by hour to CSV: %v", err)
	} else {
		fmt.Println("Highest pollutants by hour saved to highest_pollutants_by_hour.csv")
	}
}

