package main

import "fmt"
import "math"

func main() {
	var vineyard, kvM, necessaryWine, workers float64
	fmt.Scan(&vineyard, &kvM, &necessaryWine, &workers)
	vineyard *= 0.40
	allWine := vineyard * kvM / 2.5
	bonusWine := allWine - necessaryWine
	wineForWorkers := bonusWine / workers

	if necessaryWine > allWine {
		fmt.Printf("It will be a tough winter! More %v liters wine needed.", math.Abs(math.Trunc(bonusWine)))
	} else {
		fmt.Printf("Good harvest this year! Total wine: %v liters.\n%v liters left -> %v liters per person.",
			math.Floor(allWine), math.Ceil(bonusWine), math.Ceil(wineForWorkers))
	}

}

/*
name   :03. Harvest
input  :650 2 175 3
output :
Good harvest this year! Total wine: 208 liters.
33 liters left -> 11 liters per person.
*/
