package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Println(strings.Repeat("* ", n))
	}
}

/*
name   :03. Square of Stars
input  :2
output :
* *
* *
*/
