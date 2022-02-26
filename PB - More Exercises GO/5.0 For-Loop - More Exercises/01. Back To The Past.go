package main

import "fmt"
import "math"

func main() {
	var money float64
	var years int
	fmt.Scan(&money, &years)
	years_old := 17
	for i := 1800; i <= years; i++ {
		years_old += 1

		if i%2 != 0 {
			money -= 12000 + float64(years_old)*50
		} else {
			money -= 12000
		}
	}

	if money >= 0 {
		fmt.Printf("Yes! He will live a carefree life and will have %.2f dollars left.", money)
	} else {
		fmt.Printf("He will need %.2f dollars to survive.", math.Abs(money))
	}

}

/*
name   :01. Back To The Past
input  :50000 1802
output :Yes! He will live a carefree life and will have 13050.00 dollars left
*/
