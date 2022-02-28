package main

import (
	"fmt"
	"os"
)

func main() {
	var start, end, magicNum, counter int
	fmt.Scan(&start, &end, &magicNum)

	for a := start; a <= end; a++ {
		for b := start; b <= end; b++ {
			counter++
			if a+b == magicNum {
				fmt.Printf("Combination N:%v (%v + %v = %v)\n", counter, a, b, a+b)
				os.Exit(0)
			}
		}
	}

	fmt.Printf("%v combinations - neither equals %v\n", counter, magicNum)

}

/*
name   :09. Sum of Two Numbers
input  :1 10 5
output :Combination N:4 (1 + 4 = 5)
*/
