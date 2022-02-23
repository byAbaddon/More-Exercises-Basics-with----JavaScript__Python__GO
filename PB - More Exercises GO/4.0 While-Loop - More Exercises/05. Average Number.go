package main

import "fmt"

func main() {
	var n, currentNum, sum int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&currentNum)
		sum += currentNum
	}

	fmt.Printf("%.2f", float64(sum)/float64(n))
}

/*
name   :05. Average Number
input  :4 3 2 4 2
output :2.75
*/
