package main

import (
	"fmt"
	"strings"
)

func main() {
	var men, women, table int
	fmt.Scan(&men, &women, &table)
	var visitedList []string

	for m := 1; m <= men; m++ {
		for w := 1; w <= women; w++ {
			if table != len(visitedList) {
				visitedList = append(visitedList, fmt.Sprintf("(%v <-> %v)", m, w))
			} else {
				break
			}
		}
	}
	fmt.Println(strings.Trim(fmt.Sprint(visitedList), "[]"))
}

/*
name   :05. Challenge the Wedding
input  :2 2 6
output :(1 <-> 1) (1 <-> 2) (2 <-> 1) (2 <-> 2)
*/
