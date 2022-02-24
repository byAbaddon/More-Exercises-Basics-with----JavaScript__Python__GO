package main

import "fmt"

func main() {
	var b1, b2, h float32
	fmt.Scan(&b1, &b2, &h)
	fmt.Printf("%.2f", (b1+b2)*h/2)
}

/*
name   :01. Trapeziod Area
input  :8 13 7
output :73.50
*/
