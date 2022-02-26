package main

import "fmt"

func main() {
	for h := 0; h < 24; h++ {
		for m := 0; m < 60; m++ {
			for s := 0; s < 60; s++ {
				fmt.Printf("%v : %v : %v\n", h, m, s)
			}
		}
	}
}

/*
name   :10. Clock 2
input  :nil
output : 0 : 0 : 0... 23 : 59 : 59
*/
