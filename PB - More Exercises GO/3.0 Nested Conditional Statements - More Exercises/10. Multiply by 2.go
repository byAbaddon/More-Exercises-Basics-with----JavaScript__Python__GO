package main

import "fmt"

func main() {
	for {
		var n float64
		fmt.Scan(&n)
		if n < 0 {
			fmt.Println("Negative number!")
			break
		}
		fmt.Printf("Result: %.2f\n", n*2)
	}
}

/*
name   :10. Multiply by 2
input  :12 43.2144 12.3 543.23 -20
output :
Result: 24.00
Result: 86.43
Result: 24.60
Result: 1086.46
Negative number!
*/
