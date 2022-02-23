package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var detergent, dishes, pots int64
	fmt.Scan(&detergent)
	detergent *= 750

	for i := 1; i <= 100; i++ {
		var token string
		fmt.Scan(&token)
		if token == "End" {
			break
		}

		current, _ := strconv.ParseInt(token, 10, 8)

		if i%3 != 0 {
			detergent -= current * 5
			dishes += current
		} else {
			detergent -= current * 15
			pots += current
		}

		if detergent < 0 {
			break
		}
	}

	if detergent >= 0 {
		fmt.Printf("Detergent was enough!\n%v dishes and %v pots were washed.\nLeftover detergent %v ml.", dishes, pots, detergent)

	} else {
		fmt.Printf("Not enough detergent, %v ml. more necessary!", math.Abs(float64(detergent)))
	}

}

/*
name   :01. Dishwasher
input  :2 53 65 55 End
output :
Detergent was enough!
118 dishes and 55 pots were washed.
Leftover detergent 85 ml.
*/
