package main

import (
	"fmt"
	"math"
)

func main() {
	var number int
	fmt.Scan(&number)
	var evenSum, oddSum, maxDiff float64
	var isEqual = true

	for i := 1; i <= number; i++ {
		var one, two float64
		fmt.Scan(&one, &two)

		if i%2 == 0 {
			evenSum = one + two
		} else {
			oddSum = one + two
		}

		if i > 1 && math.Abs(evenSum-oddSum) > maxDiff {
			maxDiff = math.Abs(evenSum - oddSum)
			isEqual = false
		}
	}

	if isEqual {
		fmt.Printf("Yes, value=%v\n", oddSum)
	} else {

		fmt.Printf("No, maxdiff=%v\n", maxDiff)
	}

}

/*
name   :08. Equal Pairs
input  :3 1 2 0 3 4 -1
output :Yes, value=3
*/
