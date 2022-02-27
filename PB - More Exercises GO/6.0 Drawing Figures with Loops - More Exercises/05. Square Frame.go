package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	//top
	fmt.Print("+", strings.Repeat(" -", n-2), " +\n")
	//middle
	for i := 0; i < n-2; i++ {
		fmt.Println("|" + strings.Repeat(" -", n-2) + " |")
	}
	//bottom
	fmt.Print("+", strings.Repeat(" -", n-2), " +\n")
}

/*
name   :05. Square Frame
input  :3
output :
+ - +
| - |
+ - +
*/
