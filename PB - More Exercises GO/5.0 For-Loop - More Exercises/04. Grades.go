package main

import "fmt"

func main() {
	var students float64
	fmt.Scan(&students)
	var top, good, middle, fail, total float64
	for i := 0; i < int(students); i++ {
		var grade float64
		fmt.Scan(&grade)

		total += grade
		if grade >= 5 {
			top++
		} else if grade >= 4.00 && grade <= 4.99 {
			good++
		} else if grade >= 3.00 && grade <= 3.99 {
			middle++
		} else if grade <= 2.99 {
			fail++
		}
	}

	fmt.Printf("Top students: %.2f%%\n", top*100/students)
	fmt.Printf("Between 4.00 and 4.99: %.2f%%\n", good*100/students)
	fmt.Printf("Between 3.00 and 3.99: %.2f%%\n", middle*100/students)
	fmt.Printf("Fail: %.2f%%\n", fail*100/students)
	fmt.Printf("Average: %.2f", total/students)
}

/*
name   :04. Grades
input  : 10 3.00 2.99 5.68 3.01 4 4 6.00 4.50 2.44 5
output :
Top students: 30.00%
Between 4.00 and 4.99: 30.00%
Between 3.00 and 3.99: 20.00%
Fail: 20.00%
Average: 4.06
*/
