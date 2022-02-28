package main

import "fmt"

func main() {
	var oneMax, twoMax, threeMax int
	fmt.Scan(&oneMax, &twoMax, &threeMax)
	for f := 2; f <= oneMax; f += 2 {
		for s := 2; s <= twoMax; s++ {
			for t := 2; t <= threeMax; t += 2 {
				if s == 2 || s == 3 || s == 5 || s == 7 {
					fmt.Println(f, s, t)
				}
			}
		}
	}
}

/*
name   :01. Unique PIN Codes
input  :3 5 5
output :
2 2 2
2 2 4
2 3 2
2 3 4
2 5 2
2 5 4
*/
