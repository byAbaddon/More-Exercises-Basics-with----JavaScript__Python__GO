package main

import (
	"fmt"
	"os"
)

func main() {
	var password string
	var codeA = 35
	var codeB = 64
	var numA, numB, max int
	fmt.Scan(&numA, &numB, &max)

	for codeA <= 55 && codeB <= 96 {
		for x := 1; x <= numA; x++ {
			for y := 1; y <= numB; y++ {
				charA := string(codeA)
				charB := string(codeB)

				password += fmt.Sprintf("%v%v%v%v%v%v|", charA, charB, x, y, charB, charA)

				max--
				codeA++
				codeB++

				if codeA > 55 {
					codeA = 35
				}

				if codeB > 96 {
					codeB = 64
				}

				if max == 0 || (x == numA && y == numB) {
					fmt.Println(password)
					os.Exit(0)

				}
			}
		}
	}

}

/*
name   :02. Letters Combinations
input  :2 3 10
output :#@11@#|$A12A$|%B13B%|&C21C&|'D22D'|(E23E(|
*/
