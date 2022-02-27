package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	if n == 1 {
		fmt.Println("*")
		os.Exit(0)
	}

	if n == 2 {
		fmt.Println("**")
		os.Exit(0)
	}

	stars := 1
	if n%2 == 0 {
		stars = 2
	}

	//top row
	for i := 0; i < 1; i++ {
		fix := int(math.Floor(float64(n-1) / 2))
		fmt.Printf("%v%v%v\n", strings.Repeat("-", fix), strings.Repeat("*", stars), strings.Repeat("-", fix))
	}

	//middle top
	if n > 4 && n%2 != 0 { //up  /odd
		for i := 1; i < n/2; i++ {
			fix := int(math.Floor(float64(n/2 - i)))
			midFix := int(math.Floor(float64(i - 1 + i)))
			fmt.Printf("%v*%v*%v\n", strings.Repeat("-", fix), strings.Repeat("-", midFix), strings.Repeat("-", fix))
		}
	} else { //even
		for i := 1; i < n/2-1; i++ {
			fix := int(math.Floor(float64(n/2 - i - 1)))
			fmt.Printf("%v*%v*%v\n", strings.Repeat("-", fix), strings.Repeat("-", i+i), strings.Repeat("-", fix))
		}
	}

	// middle line
	for i := 0; i < 1; i++ {
		fmt.Printf("*%v*\n", strings.Repeat("-", n-2))
	}

	//under middle
	if n > 4 && n%2 != 0 { //odd
		for i := n/2 - 1; i > 0; i-- {
			fix := int(math.Floor(float64(n/2 - i)))
			fmt.Printf("%v*%v*%v\n", strings.Repeat("-", fix), strings.Repeat("-", i+i-1), strings.Repeat("-", fix))
		}

	} else { //even
		for i := n/2 - 1; i > 1; i-- {
			fix := int(math.Floor(float64(n/2 - i)))
			fmt.Printf("%v*%v*%v\n", strings.Repeat("-", fix), strings.Repeat("-", i+i-2), strings.Repeat("-", fix))
		}
	}

	//buttom row
	for i := 0; i < 1; i++ {
		fix := int(math.Floor(float64(n-1) / 2))
		fmt.Printf("%v%v%v\n", strings.Repeat("-", fix), strings.Repeat("*", stars), strings.Repeat("-", fix))
	}

}

/*
name   :10. Diamond
input  :6
output :
--**--
-*--*-
*----*
-*--*-
--**--
*/
