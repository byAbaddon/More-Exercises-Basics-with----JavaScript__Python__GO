package main

import "fmt"

func main() {

	for m := 0; m < 24; m++ {
		for s := 0; s < 60; s++ {
			fmt.Printf("%v : %v\n", m, s)
		}
	}

}

/*
name   :09 Clock
input  :nil
output : 0 : 0 ... 23 : 59
*/
