package main

import "fmt"

func main() {
	var chrysanthemums, roses, tulips float64
	var season, isCelebrateDay string
	fmt.Scan(&chrysanthemums, &roses, &tulips, &season, &isCelebrateDay)
	var chrysanthemumTotal, roseTotal, tulipTotal float64
	if season == "Spring" || season == "Summer" {
		chrysanthemumTotal = chrysanthemums * 2
		roseTotal = roses * 4.1
		tulipTotal = tulips * 2.5
	} else if season == "Autumn" || season == "Winter" {
		chrysanthemumTotal = chrysanthemums * 3.75
		roseTotal = roses * 4.5
		tulipTotal = tulips * 4.15
	}

	price := chrysanthemumTotal + roseTotal + tulipTotal

	if isCelebrateDay == "Y" {
		price *= 1.15
	}

	if season == "Spring" && tulips > 7 {
		price *= 0.95
	}

	if season == "Winter" && roses >= 10 {
		price *= 0.90
	}

	if chrysanthemums+roses+tulips > 20 {
		price *= 0.8
	}

	fmt.Printf("%.2f", price+2)
}

/*
name   :03. Flowers
input  :2 4 8 Spring Y
output :46.14
*/
