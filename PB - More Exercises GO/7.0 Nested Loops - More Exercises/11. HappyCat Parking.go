package main

import "fmt"

func main() {
	var days, hours, currentTax, totalSum float64
	fmt.Scan(&days, &hours)

	for day := 1; day <= int(days); day++ {
		for hour := 1; hour <= int(hours); hour++ {
			if day%2 == 0 && hour%2 != 0 {
				currentTax += 2.50
			} else if day%2 != 0 && hour%2 == 0 {
				currentTax += 1.25
			} else {
				currentTax += 1.00
			}
		}
		fmt.Printf("Day: %v - %.2f leva\n", day, currentTax)
		totalSum += currentTax
		currentTax = 0
	}

	fmt.Printf("Total: %.2f leva\n", totalSum)
}

/*
name   : pacman -Qi qt5-base qt5ct
input  :2 5
output :
Day: 1 - 5.50 leva
Day: 2 - 9.50 leva
Total: 15.00 leva
*/
