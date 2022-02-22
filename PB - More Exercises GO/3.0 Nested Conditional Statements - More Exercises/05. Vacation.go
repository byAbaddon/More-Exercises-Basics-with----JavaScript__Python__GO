package main

import "fmt"

func main() {
	var budget float64
	var season, location, placeType string
	fmt.Scan(&budget, &season)

	if budget <= 1000 {
		placeType = "Camp"
		if season == "Summer" {
			location = "Alaska"
			budget *= 0.65
		} else {
			location = "Morocco"
			budget *= 0.45
		}
	} else if 1000 < budget && budget <= 3000 {
		placeType = "Hut"
		if season == "Summer" {
			location = "Alaska"
			budget *= 0.80
		} else {
			location = "Morocco"
			budget *= 0.60
		}
	} else {
		placeType = "Hotel"
		if season == "Summer" {
			location = "Alaska"
			budget *= 0.90
		} else {
			location = "Morocco"
			budget *= 0.90
		}
	}

	fmt.Printf("%v - %v - %.2f", location, placeType, budget)
}

/*
name   :05. Vacation
input  :800 Summer
output :Alaska - Camp - 520.00
*/
