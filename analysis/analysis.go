package analysis

import (
	"air-quality-analyzer/models"
)

// CalculateAverage calculates the average values of pollutants from sensor readings
func CalculateAverage(readings []models.AirQualityReading) map[string]float64 {
	var totalPM25, totalCO2 float64
	count := float64(len(readings))

	for _, reading := range readings {
		totalPM25 += reading.PM25
		totalCO2 += reading.CO2
	}

	averages := map[string]float64{
		"pm25": totalPM25 / count,
		"co2":  totalCO2 / count,
	}

	return averages
}

// FindHighestPollutantByHour finds the pollutant with the highest average value for each hour
func FindHighestPollutantByHour(readings []models.AirQualityReading) (map[int]string, map[int]map[string]float64) {

	// เก็บข้อมูลผลรวมของสารมลพิษตามชั่วโมง
	hourlyData := make(map[int]map[string]float64)
	// เก็บจำนวน readings ของสารมลพิษในแต่ละชั่วโมง
	hourlyCount := make(map[int]map[string]int)

	// วนลูปผ่าน readings และจัดกลุ่มตามชั่วโมง
	for _, reading := range readings {
		hour := reading.Timestamp.Hour()

		// ตรวจสอบว่ามีข้อมูลสารมลพิษในชั่วโมงนั้นหรือไม่ ถ้าไม่มีให้สร้างใหม่
		if hourlyData[hour] == nil {
			hourlyData[hour] = make(map[string]float64)
			hourlyCount[hour] = make(map[string]int)
		}

		// เก็บค่าสารมลพิษ pm25 และ co2
		hourlyData[hour]["pm25"] += reading.PM25
		hourlyData[hour]["co2"] += reading.CO2
		hourlyCount[hour]["pm25"]++
		hourlyCount[hour]["co2"]++
	}

	
	highestPollutantByHour := make(map[int]string)
	
	averagePollutants := make(map[int]map[string]float64)

	// คำนวณหาสารมลพิษที่มีค่าเฉลี่ยสูงสุดในแต่ละชั่วโมง
	for hour, pollutants := range hourlyData {
		var highestPollutant string
		var highestAverage float64

		averagePollutants[hour] = make(map[string]float64) // สร้างแผนที่สำหรับเก็บค่าเฉลี่ยของสารมลพิษในชั่วโมงนี้

		for pollutant, totalValue := range pollutants {
			count := float64(hourlyCount[hour][pollutant])
			if count > 0 {
				average := totalValue / count
				averagePollutants[hour][pollutant] = average // บันทึกค่าเฉลี่ยของสารมลพิษในชั่วโมงนี้

				// ตรวจสอบว่าค่าเฉลี่ยสูงสุดหรือไม่
				if average > highestAverage {
					highestAverage = average
					highestPollutant = pollutant
				}
			}
		}

		// บันทึกผลลัพธ์สำหรับแต่ละชั่วโมง
		if highestPollutant != "" {
			highestPollutantByHour[hour] = highestPollutant
		}
	}

	return highestPollutantByHour, averagePollutants
}

