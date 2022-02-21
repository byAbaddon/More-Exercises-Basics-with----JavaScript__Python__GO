package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64
	fmt.Scan(&n)
	face := math.Pi * n * n
	prm :=  math.Pi * 2 * n
	fmt.Printf("%.2f\n%.2f", face, prm)
}

/*
name   :08. Circle Area and Perimeter
input  :3
output :28.27 / 18.85
*/
