package main

import "fmt"

func main() {
	var budget float64
	var season, classType, carType string
	fmt.Scan(&budget, &season)

	if budget <= 100 {
		classType = "Economy class"
		if season == "Summer" {
			budget *= 0.35
			carType = "Cabrio"
		} else {
			budget *= 0.65
			carType = "Jeep"
		}
	} else if 100 < budget && budget <= 500 {
		classType = "Compact class"
		if season == "Summer" {
			budget *= 0.45
			carType = "Cabrio"
		} else {
			budget *= 0.80
			carType = "Jeep"
		}
	} else {
		classType = "Luxury class"
		budget *= 0.90
		carType = "Jeep"
	}

	fmt.Printf("%v\n%v - %.2f", classType, carType, budget)
}

/*
name   :04. Car To Go
input  :450 Summer
output :
Compact class
Cabrio - 202.50
*/
