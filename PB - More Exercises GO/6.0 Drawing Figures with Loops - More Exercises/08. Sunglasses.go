package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var el string
	var n int
	fmt.Scan(&n)

	//top
	fmt.Printf("%v%v%v\n", strings.Repeat("*", 2*n), strings.Repeat(" ", n), strings.Repeat("*", 2*n))

	//middle
	for i := 0; i < n-2; i++ {
		if i == int(math.Ceil(float64((n-1)/2-1))) {
			el = strings.Repeat("|", n)
		} else {
			el = strings.Repeat(" ", n)
		}

		fmt.Printf("*%v*%v*%v*\n", strings.Repeat("/", n+(n-2)), el, strings.Repeat("/", n+(n-2)))

	}

	//bottom
	fmt.Printf("%v%v%v\n", strings.Repeat("*", 2*n), strings.Repeat(" ", n), strings.Repeat("*", 2*n))
}

/*
name   :08. Sunglasses
input  :3
output :
******   ******
*/ ///*|||*////*
/******   ******
 */
