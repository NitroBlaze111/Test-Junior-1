package main

import (
	"air-quality-analyzer/analysis"
	"air-quality-analyzer/persistence"
	"air-quality-analyzer/processing"
	"fmt"
	"io/ioutil"
)

func main() {
	// โหลดข้อมูลจาก mockdata.json
	data, _ := ioutil.ReadFile("data/mockdata.json")

	// Parse JSON เป็น AirQualityReading structs
	readings, err := processing.ParseReadings(data)
	if err != nil {
		fmt.Println("Error parsing data:", err)
		return
	}

	// คำนวณค่าเฉลี่ยของแต่ละสารมลพิษ
	averages := analysis.CalculateAverage(readings)
	fmt.Println("Average values:", averages)

	// หาสารมลพิษที่มีค่าเฉลี่ยสูงสุดในแต่ละชั่วโมง
	highestPollutantByHour := analysis.FindHighestPollutantByHour(readings)
	fmt.Println("Highest Pollutant By Hour:", highestPollutantByHour)

	// บันทึกค่าเฉลี่ยลงใน CSV
	err = persistence.SaveAverageToCSV(averages, "average_values.csv")
	if err != nil {
		fmt.Println("Error saving average values to CSV:", err)
	} else {
		fmt.Println("Average values saved to average_values.csv")
	}

	// บันทึกสารมลพิษที่มีค่าสูงสุดในแต่ละชั่วโมงลงใน CSV
	err = persistence.SaveToCSV(highestPollutantByHour, "highest_pollutants_by_hour.csv")
	if err != nil {
		fmt.Println("Error saving highest pollutants by hour to CSV:", err)
	} else {
		fmt.Println("Highest pollutants by hour saved to highest_pollutants_by_hour.csv")
	}
}
