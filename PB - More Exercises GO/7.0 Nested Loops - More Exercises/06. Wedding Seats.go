package main

import "fmt"

func main() {
	var sectors string
	var rowCount, seatsCount, evenSeats, counter int
	fmt.Scan(&sectors, &rowCount, &seatsCount)

	for sector := 'A'; sector <= 'Z'; sector++ {
		rowCount++
		for row := 1; row < rowCount; row++ {
			if row%2 == 0 {
				evenSeats = 2
			} else {
				evenSeats = 0
			}

			for seats := 1; seats < (seatsCount + evenSeats + 1); seats++ {
				counter++
				fmt.Printf("%v%v%v\n", string(sector), row, string(96+seats))
			}
		}
		if string(sector) == sectors {
			fmt.Println(counter)
			break
		}
	}

}

/*
name   :06. Wedding Seats
input  :B 3 2
output :
A1a
A1b
A2a
A2b
A2c
A2d
A3a
A3b
B1a
B1b
B2a
B2b
B2c
B2d
B3a
B3b
B4a
B4b
B4c
B4d
20
*/
