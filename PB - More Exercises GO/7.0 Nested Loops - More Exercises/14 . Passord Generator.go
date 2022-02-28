package main

import "fmt"
import "math"

func main() {
	var num, word int
	fmt.Scan(&num, &word)
	const str = "abcdefghijklmnopqrstuvwxyz"
	var numList string

	for a := 1; a < num; a++ {
		for b := 1; b < num; b++ {
			for c := 0; c < word; c++ {
				for d := 0; d < word; d++ {
					for e := math.Max(float64(a), float64(b)) + 1; int(e) <= num; e++ {
						numList += fmt.Sprintf("%v%v%v%v%v ", a, b, string(str[c]), string(str[d]), e)
					}
				}
			}
		}
	}

	fmt.Println(numList)
}

/*
name   :passwordGenerator
input  : 3, 1
output :11aa2 11aa3 12aa3 21aa3 22aa3
*/
