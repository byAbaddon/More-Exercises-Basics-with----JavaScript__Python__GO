package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	//first part
	for i := 1; i < n; i++ {
		fmt.Println(strings.Repeat(" ", n-i-1), strings.Repeat("* ", i))

	}

	//middle
	fmt.Println(strings.Repeat("* ", n))

	//end part
	for i := 1; i < n; i++ {
		fmt.Println(strings.Repeat(" ", i-1), strings.Repeat("* ", n-i))
	}
}

/*
name   :06. Rhombus of Stars
input  :3
output :
   *
  * *
 * * *
  * *
   *
*/
