package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var result string

	for a := 1; a <= 9; a++ {
		for b := 1; b <= 9; b++ {
			for c := 1; c <= 9; c++ {
				for d := 1; d <= 9; d++ {
					if a+b == c+d && n%(c+d) == 0 {
						result += fmt.Sprintf("%v%v%v%v ", a, b, c, d)
					}
				}
			}
		}
	}

	fmt.Println(result)
}

/*
name   :03. Lucky Numbers
input  :3
output :1212 1221 2112 2121
*/
