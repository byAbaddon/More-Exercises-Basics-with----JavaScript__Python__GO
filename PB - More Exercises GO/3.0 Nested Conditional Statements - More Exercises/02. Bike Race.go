package main

import "fmt"

func main() {
	var numJuniorBiker, numSeniorBiker float64
	var routeType string
	fmt.Scan(&numJuniorBiker, &numSeniorBiker, &routeType)
	groupMap := map[string]map[string]float64{
		"juniors": {"trail": 5.50, "cross-country": 8.00, "downhill": 12.25, "road": 20.00},
		"seniors": {"trail": 7.00, "cross-country": 9.50, "downhill": 13.75, "road": 21.50},
	}

	juniorsMoney := groupMap["juniors"][routeType] * numJuniorBiker
	seniorsMoney := groupMap["seniors"][routeType] * numSeniorBiker
	allBikersMoney := juniorsMoney + seniorsMoney
	bikerCount := numJuniorBiker + numSeniorBiker

	if routeType == "cross-country" && bikerCount >= 50 {
		allBikersMoney = allBikersMoney * 0.75
	}

	allBikersMoney = allBikersMoney * 0.95

	fmt.Printf("%.2f", allBikersMoney+0.001)
}

/*
name   :02. Bike Race
input  :10 20 trail
output :185.25

input: 20 25 cross-country
output:377.63
*/
