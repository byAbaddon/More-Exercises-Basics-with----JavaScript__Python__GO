package main

import "fmt"

func main() {
	var one, two, ignore string
	fmt.Scan(&one, &two, &ignore)

	charsLegend := map[string]int{}
	for i := 'a'; i <= 'z'; i++ {
		charsLegend[string(i)] = int(i)
	}


	start := charsLegend[one]
	end := charsLegend[two]
	check := charsLegend[ignore]
	counter := 0

	for i := start; i <= end; i++ {
		for j := start; j <= end; j++ {
			for k := start; k <= end; k++ {
				if i != check && j != check && k != check {
					fmt.Printf(string(i) + string(j) + string(k) + " ")
					counter++
				}
			}
		}
	}

	fmt.Print(counter)
}

/*
name   :02. Letters Combinations
input  :a c b
output :aaa aac aca acc caa cac cca ccc 8
*/
