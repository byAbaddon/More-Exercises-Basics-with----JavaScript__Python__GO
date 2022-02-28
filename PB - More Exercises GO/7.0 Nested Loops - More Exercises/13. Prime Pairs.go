
package main

import "fmt"
import "math"

func main() {
	var firstPair, secondPair, firstDiff, secondDiff int
	fmt.Scan(&firstPair, &secondPair, &firstDiff, &secondDiff)

	for first := firstPair; first <= firstPair+firstDiff; first++ {
		for second := secondPair; second <= secondPair+secondDiff; second++ {

			firstSqrt := math.Floor(math.Sqrt(float64(first)))
			secondSqrt := math.Floor(math.Sqrt(float64(second)))
			isPrime1 := true
			isPrime2 := true

			for first1 := 2; first1 <= int(firstSqrt); first1++ {
				if first%first1 == 0 {
					isPrime1 = false
				}
			}

			for second2 := 2; second2 <= int(secondSqrt); second2++ {
				if second%second2 == 0 {
					isPrime2 = false
				}
			}

			if isPrime1 == true && isPrime2 == true {
				fmt.Printf("%v%v\n", first, second)

			}

		}
	}

}


/*
name   :13. Prime Pairs
input  :10 20 5 5
output :
*/

//------------------------------(2)----------------Fucking Judge

// package main

// import "fmt"

// func main() {
// 	var fStart, sStart, fEnd, sEnd int
// 	fmt.Scan(&fStart, &sStart, &fEnd, &sEnd)
// 	fEnd += fStart
// 	sEnd += sStart
// 	primeList := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}

// 	for i := fStart; i <= fEnd; i++ {
// 		for j := sStart; j <= sEnd; j++ {
// 			for _, v := range primeList {
// 				if v == i {
// 					for _, v := range primeList {
// 						if v == j {
// 							print(i, j, "\n")
// 							break
// 						}
// 					}
// 				}

// 			}

// 		}
// 	}

// }
