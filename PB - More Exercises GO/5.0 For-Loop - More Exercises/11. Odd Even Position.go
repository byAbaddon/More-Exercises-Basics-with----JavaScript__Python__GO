package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var oddSum, oddMin, oddMax = 0.0, 10000000000000000.0, -999999999999999.0
	var evenSum, evenMin, evenMax = 0.0, 10000000000000000.0, -999999999999999.0
	var oddMinR, oddMaxR, evenMinR, evenMaxR string

	for i := 1; i <= n; i++ {
		var el float64
		fmt.Scan(&el)

		if i%2 != 0 {
			oddSum += el
			if el > oddMax {
				oddMax = el
			}
			if el < oddMin {
				oddMin = el
			}

		} else {
			evenSum += el
			if el > evenMax {
				evenMax = el
			}
			if el < evenMin {
				evenMin = el
			}
		}

	}

	if oddMin == 10000000000000000 {
		oddMinR = "No"
	} else {

		oddMinR = fmt.Sprintf("%.2f", float64(oddMin))

	}

	if oddMax == -999999999999999 {
		oddMaxR = "No"
	} else {
		oddMaxR = fmt.Sprintf("%.2f", float64(oddMax))
	}

	if evenMin == 10000000000000000 {
		evenMinR = "No"
	} else {
		evenMinR = fmt.Sprintf("%.2f", float64(evenMin))
	}

	if evenMax == -999999999999999{
		evenMaxR = "No"
	} else {
		evenMaxR = fmt.Sprintf("%.2f", float64(evenMax))
	}

	fmt.Printf("OddSum=%.2f,\nOddMin=%v,\nOddMax=%v,\n", float64(oddSum), oddMinR, oddMaxR)

	fmt.Printf("EvenSum=%.2f,\nEvenMin=%v,\nEvenMax=%v\n", float64(evenSum), evenMinR, evenMaxR)
}

/*
name   :11. Odd / Even Position
input  :
6
2
3
5
4
2
1

output :
OddSum=9.00, OddMin=2.00, OddMax=5.00, EvenSum=8.00, EvenMin=1.00, EvenMax=4.00
*/
