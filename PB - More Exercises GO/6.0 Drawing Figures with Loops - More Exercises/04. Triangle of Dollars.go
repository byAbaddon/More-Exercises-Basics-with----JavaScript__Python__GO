package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Println(strings.Repeat("$ ", i))
	}
}

/*
name   :04. Triangle of Dollars
input  :3
output :
$
$ $
$ $ $
*/
