package main

import (
	"fmt"
	"math"
)

func main() {
	var hours, day, workers float64
	fmt.Scan(&hours, &day, &workers)

	remnantTime := math.Trunc(day * 0.90 * 8)
	extraWorkTime := workers * (2 * day)
	totalTime := remnantTime + extraWorkTime

	if totalTime >= hours {
		fmt.Printf("Yes!%.0f hours left.", totalTime-hours)
	} else {
		fmt.Printf("Not enough time!%.0f hours needed.", hours-totalTime)
	}
}

/*
name   :05. Firm
input  :90 7 3
output :Yes!2 hours left.
*/
