package main

import "fmt"

func main() {
	var temp float64
	fmt.Scan(&temp)

	if 5 <= temp && temp <= 11.9 {
		fmt.Println("Cold")
	} else if 12 <= temp && temp <= 14.9 {
		fmt.Println("Cool")
	} else if 15 <= temp && temp <= 20 {
		fmt.Println("Mild")
	} else if 20.1 <= temp && temp <= 25.9 {
		fmt.Println("Warm")
	} else if 26 <= temp && temp <= 35 {
		fmt.Println("Hot")
	} else {
		fmt.Println("unknown")
	}
}

/*
name   :10. Weather Forecast - Part 2
input  :8
output :Cold
*/

// function weatherForecastTwo(temp) {
//     if (5 <= temp && temp <= 11.9) return 'Cold'
//     if (12 <= temp && temp <= 14.9) return 'Cool'
//     if (15 <= temp && temp <= 20) return 'Mild'
//     if (20.1 <= temp && temp <= 25.9) return 'Warm'
//     if (26 <= temp && temp <= 35) return 'Hot'
//     return 'unknown'
// }

// console.log(weatherForecastTwo(8)) 'Cold'
