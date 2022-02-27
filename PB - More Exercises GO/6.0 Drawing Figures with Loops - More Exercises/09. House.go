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
	if n == 2 {
		fmt.Println("**\n" + "||")
		os.Exit(0)
	}

	//top
	if n%2 == 0 {
		for i := 1; i < (n / 2); i++ { //is even
			fmt.Printf("%v%v%v\n", strings.Repeat("-", n/2-i), strings.Repeat("*", i*2), strings.Repeat("-", n/2-i))
		}
	} else {
		for i := 1; i < n/2+1; i++ { //is odd
			fix := math.Floor(float64(n/2 - i + 1))
			fmt.Printf("%v%v%v\n", strings.Repeat("-", int(fix)), strings.Repeat("*", i*2-1), strings.Repeat("-", int(fix)))
		}
	}

	//middle part
	fmt.Println(strings.Repeat("*", n))

	//bottom
	if n%2 == 0 {
		for i := 0; i < n/2; i++ {
			fmt.Printf("|%v|\n", strings.Repeat("*", n-2))
		}
	}

	if n%2 != 0 {
		for i := 0; i < n/2; i++ {
			fmt.Printf("|%v|\n", strings.Repeat("*", n-2))
		}
	}

}

/*
name   :09. House
input  :3
output :
-*-
***
|*|
*/
