package main

import (
	"fmt"
	"math"
)

func main() {
	var countMgn, countZum, countRoses, countCactus, giftPrice float64
	fmt.Scan(&countMgn, &countZum, &countRoses, &countCactus, &giftPrice)

	tax := 0.95
	totalSum := (countMgn*3.25 + countZum*4 + countRoses*3.5 + countCactus*8) * tax
	if totalSum < giftPrice {
		fmt.Printf("She will have to borrow %v leva.", math.Ceil(giftPrice-totalSum))
	} else {
		fmt.Printf("She is left with %v leva.", math.Floor(totalSum-giftPrice))
	}
}

/*
name   :07. Flower Shop
input  :2 3 5 1 50
output :She will have to borrow 9 leva.
*/
