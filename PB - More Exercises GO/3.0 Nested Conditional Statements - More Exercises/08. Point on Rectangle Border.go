package main

import "fmt"

func main() {
	var x1, y1, x2, y2, x, y float64
	fmt.Scan(&x1, &y1, &x2, &y2, &x, &y)

	if (x == x1 && y1 <= y && y <= y2) || (x == x2 && y1 <= y && y <= y2) ||
		(y == y1 && x1 <= x && x <= x2) || (y == y2 && x1 <= x && x <= x2) {
		fmt.Println("Border")
	} else if y == y1 && y == y2 && x1 < x && x < x2 {
		fmt.Println("Border")
	} else {
		fmt.Println("Inside / Outside")
	}

}

/*
name   :08. Point on Rectangle Border
input  :2 -3 12 3 1 3
output :Inside / Outside
*/
