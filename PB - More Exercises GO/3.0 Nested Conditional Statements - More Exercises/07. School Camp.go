package main

import "fmt"

func main() {
	var season, gender string
	var countStudents, countNights float64
	fmt.Scan(&season, &gender, &countStudents, &countNights)
	holidayType := map[string]map[string]string{
		"girls": {"Winter": "Gymnastics", "Spring": "Athletics", "Summer": "Volleyball"},
		"boys":  {"Winter": "Judo", "Spring": "Tennis", "Summer": "Football"},
		"mixed": {"Winter": "Ski", "Spring": "Cycling", "Summer": "Swimming"},
	}

	priceType := map[string]map[string]float64{
		"girls": {"Winter": 9.60, "Spring": 7.20, "Summer": 15.00},
		"boys":  {"Winter": 9.60, "Spring": 7.20, "Summer": 15.00},
		"mixed": {"Winter": 10.00, "Spring": 9.50, "Summer": 20.00},
	}

	sportType := holidayType[gender][season]
	totalPrice := priceType[gender][season] * countStudents * countNights

	if 10 <= countStudents && countStudents < 20 {
		totalPrice *= 0.95
	} else if 20 <= countStudents && countStudents < 50 {
		totalPrice *= 0.85
	} else if countStudents >= 50 {
		totalPrice *= 0.50
	}

	fmt.Printf("%v %.2f lv.", sportType, totalPrice)
}

/*
name   :07. School Camp
input  :Spring girls 20 7
output :Athletics 856.80 lv.
*/
