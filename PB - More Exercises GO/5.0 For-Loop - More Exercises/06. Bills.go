package main

import "fmt"

func main() {
	var monthsCount, allElectricity float64
	fmt.Scan(&monthsCount)
	for i := 0; i < int(monthsCount); i++ {
		var current float64
		fmt.Scan(&current)
		allElectricity += current
	}
	water := 20 * monthsCount
	internet := 15 * monthsCount
	other := allElectricity*1.20 + water*1.20 + internet*1.2
	average := (allElectricity + water + internet + other) / monthsCount

	fmt.Printf("Electricity: %.2f lv\n", allElectricity)
	fmt.Printf("Water: %.2f lv\n", water)
	fmt.Printf("Internet: %.2f lv\n", internet)
	fmt.Printf("Other: %.2f lv\n", other)
	fmt.Printf("Average: %.2f lv\n", average)
}

/*
name   :06. Bills
input  :5 68.63 89.25 132.53 93.53 63.22
output :
Electricity: 447.16 lv
Water: 100.00 lv
Internet: 75.00 lv
Other: 746.59 lv
Average: 273.75 lv
*/
