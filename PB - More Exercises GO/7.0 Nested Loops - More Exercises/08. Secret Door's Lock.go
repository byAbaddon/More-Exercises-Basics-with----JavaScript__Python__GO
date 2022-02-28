package main

import "fmt"

func main() {
	var hundredNum, doubleNum, singleNum int
	fmt.Scan(&hundredNum, &doubleNum, &singleNum)
	for h := 2; h <= hundredNum; h += 2 {
		for d := 2; d <= doubleNum; d++ {
			for s := 2; s <= singleNum; s += 2 {
				if d == 2 || d == 3 || d == 5 || d == 7 {
					fmt.Println(h, d, s)
				}
			}
		}
	}
}

/*
name   :08. Secret Door's Lock
input  :2 5 5
output :
2 2 2
2 2 4
2 3 2
2 3 4
2 5 2
2 5 4
*/
