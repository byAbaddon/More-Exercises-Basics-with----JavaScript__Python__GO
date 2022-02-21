package main

import "fmt"

func main() {
	var km float64
	var day string
	fmt.Scan(&km, &day)
	price := 0.0

	if km < 20 {
		if day == "day" {
			price = 0.79*km + 0.70
		} else {
			price = 0.90*km + 0.70
		}
	} else if 20 <= km && km < 100 {
		price = km * 0.09
	} else {
		price = km * 0.06
	}
	fmt.Printf("%.2f", price)
}

/*
name   :04. Transport Price
input  :7 night
output :7.00
*/
