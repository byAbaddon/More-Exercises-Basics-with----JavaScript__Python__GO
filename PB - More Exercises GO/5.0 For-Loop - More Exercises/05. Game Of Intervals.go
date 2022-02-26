package main

import "fmt"

func main() {
	var move, points float64
	fmt.Scan(&move)

	percentObj := map[string]float64{"From0": 0.0, "From10": 0.0, "From20": 0.0, "From30": 0.0, "From40:": 0.0, "Invalid numbers:": 0.0}

	for i := 0; i < int(move); i++ {
		var currentNum float64
		fmt.Scan(&currentNum)

		if 0 <= currentNum && currentNum <= 9 {
			points += currentNum * 0.20
			percentObj["From0"] += 1
		} else if 10 <= currentNum && currentNum <= 19 {
			points += currentNum * 0.30
			percentObj["From10"] += 1
		} else if 20 <= currentNum && currentNum <= 29 {
			points += currentNum * 0.40
			percentObj["From20"] += 1
		} else if 30 <= currentNum && currentNum <= 39 {
			points += 50
			percentObj["From30"] += 1
		} else if 40 <= currentNum && currentNum <= 50 {
			points += 100
			percentObj["From40"] += 1
		} else {
			points /= 2
			percentObj["Invalid numbers:"] += 1
		}
	}

	fmt.Printf("%.2f\n", points)
	fmt.Printf("From 0 to 9: %.2f%%\n", percentObj["From0"] / move * 100)
	fmt.Printf("From 10 to 19: %.2f%%\n", percentObj["From10"] / move * 100)
	fmt.Printf("From 20 to 29: %.2f%%\n", percentObj["From20"] / move * 100)
	fmt.Printf("From 30 to 39: %.2f%%\n", percentObj["From30"] / move * 100)
	fmt.Printf("From 40 to 50: %.2f%%\n", percentObj["From40"] / move * 100)
	fmt.Printf("Invalid numbers: %.2f%%\n", percentObj["Invalid numbers:"] / move * 100)

}

/*
name   :05. GameOfIntervals
input  :
10
43
57
-12
23
12
0
50
40
30
20

output :
295.80
From 0 to 9: 10.00%
From 10 to 19: 10.00%
From 20 to 29: 20.00%
From 30 to 39: 10.00%
From 40 to 50: 30.00%
Invalid numbers: 20.00%

*/
