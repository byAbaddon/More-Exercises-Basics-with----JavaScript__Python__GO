package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	//head
	for i := 0; i < 1; i++ {
		fmt.Println(strings.Repeat(" ", n) + " | " + strings.Repeat(" ", n))
	}
	//body
	for i := 1; i <= n; i++ {
		fmt.Printf("%v%v | %v\n", strings.Repeat(" ", n-i), strings.Repeat("*", i), strings.Repeat("*", i))
	}

}

/*
name   :07. Christmas Tree
input  :3
output :
    |    
  * | *
 ** | **
*** | ***
*/
