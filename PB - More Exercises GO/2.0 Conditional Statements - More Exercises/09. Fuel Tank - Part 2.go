package main

import "fmt"

func main() {
	var quantity float64
	var typeFuel, clubCard string
	fmt.Scan(&typeFuel, &quantity, &clubCard)

	fuelPriceDict := map[string]float64{"Gasoline": 2.22, "Diesel": 2.33, "Gas": 0.93}
	discount := map[string]float64{"Gasoline": 0.18, "Diesel": 0.12, "Gas": 0.08}
	subtotalPrice := 0.0
	totalPrice := 0.0

	if clubCard == "Yes" {
		subtotalPrice = fuelPriceDict[typeFuel] - discount[typeFuel]
		totalPrice = subtotalPrice * quantity
	} else {
		totalPrice = fuelPriceDict[typeFuel] * quantity
	}

	if 20 <= quantity && quantity <= 25 {
		totalPrice *= 0.92
	} else if quantity > 25 {
		totalPrice *= 0.90
	}

	fmt.Printf("%.2f lv.", totalPrice)

}

/*
name   :09. Fuel Tank - Part 2
input  :Gas 30 Yes
output :22.95 lv.
*/
