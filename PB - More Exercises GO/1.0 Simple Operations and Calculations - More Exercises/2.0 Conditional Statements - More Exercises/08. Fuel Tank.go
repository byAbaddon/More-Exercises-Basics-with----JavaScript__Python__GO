package main

import (
	"fmt"
	"strings"
)

func main() {
	var typeFuel string
	var quantity float64
	fmt.Scan(&typeFuel, &quantity)

	if typeFuel == "Diesel" || typeFuel == "Gasoline" || typeFuel == "Gas" {
		if quantity >= 25 {
			fmt.Printf("You have enough %v.", strings.ToLower(typeFuel))
		} else if quantity < 25 {
			fmt.Printf("Fill your tank with %v!", strings.ToLower(typeFuel))
		}
	} else {
		fmt.Println("Invalid fuel!")
	}
}

/*
name   :08. Fuel Tank
input  :Diesel 10
output :Fill your tank with diesel!
*/
