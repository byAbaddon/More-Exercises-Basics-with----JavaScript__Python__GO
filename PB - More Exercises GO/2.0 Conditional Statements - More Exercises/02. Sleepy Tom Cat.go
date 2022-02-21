package main

import (
	"fmt"
	"math"
)

func main() {
	var days float64
	fmt.Scan(&days)
	gameTomMin := 30000.00
	year := (365 - days)
	yearToMin := year*63 + days*127
	subtotal := math.Abs(gameTomMin - yearToMin)
	hours := math.Trunc(subtotal / 60)
	min := math.Mod(subtotal, 60)

	if gameTomMin >= yearToMin {
		fmt.Printf("Tom sleeps well\n%v hours and %v minutes less for play", hours, min)
	} else {
		fmt.Printf("Tom will run away\n%v hours and %v minutes more for play", hours, min)
	}

}

/*
name   :02. Sleepy Tom Cat
input  :20
output :
Tom sleeps well
95 hours and 25 minutes less for play
*/
