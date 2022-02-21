package main

import "fmt"

func main() {
	var volume, pipeOne, pipeTwo, hours float64
	fmt.Scan(&volume, &pipeOne, &pipeTwo, &hours)
	pipeOne *= hours
	pipeTwo *= hours
	pool := pipeOne + pipeTwo
	poolPercent := (pool / volume) * 100
	pipeOnePercent := (pipeOne / pool) * 100
	pipeTwoPercent := (pipeTwo / pool) * 100

	if pool <= volume {
		fmt.Printf("The pool is %.2f%% full. Pipe 1: %.2f%%. Pipe 2: %.2f%%.", poolPercent, pipeOnePercent, pipeTwoPercent)
	} else {
		fmt.Printf("For %.2f hours the pool overflows with %.2f liters.", hours, pool-volume)
	}

}

/*
 name   :01. Pipes In Pool
 input  :
1000
100
120
3
output : The pool is 66.00% full. Pipe 1: 45.45%. Pipe 2: 54.55%.
*/
