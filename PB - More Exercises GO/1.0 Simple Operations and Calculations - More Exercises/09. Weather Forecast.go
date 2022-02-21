package main

import "fmt"

func main() {
	var weather string
	fmt.Scan(&weather)
	if weather == "sunny" {
		fmt.Println("It's warm outside!")
	} else if weather == "cloudy" {
		fmt.Println("It's cold outside!")
	} else {
		fmt.Println("It's cold outside!")
	}
}

/*
name   :09. Weather Forecast
input  :sunny
output :It's warm outside!
*/
