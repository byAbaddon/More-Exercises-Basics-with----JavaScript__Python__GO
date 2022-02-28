package main

import "fmt"

func main() {
	var start, end int
	fmt.Scan(&start, &end)
	var result string

	for a := start; a <= end; a++ {
		for b := start; b <= end; b++ {
			for c := start; c <= end; c++ {
				for d := start; d <= end; d++ {
					if a%2 == 0 && d%2 != 0 || a%2 != 0 && d%2 == 0 {
						if a > d {
							if (b+c)%2 == 0 {
								result += fmt.Sprintf("%v%v%v%v ", a, b, c, d)
							}
						}
					}
				}
			}
		}
	}

	fmt.Println(result)
}

/*
name   :04. Car Number
input  :2 3
output :3222 3332
*/
