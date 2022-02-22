package main

import (
	"fmt"
	"math"
)

func main() {
	var ticket string
	var budget, people float64
	fmt.Scan(&budget, &ticket, &people)
	ticketsType := map[string]float64{"VIP": 499.99, "Normal": 249.99}

	if 1 <= people && people <= 4 {
		budget *= 0.25
	} else if 5 <= people && people <= 9 {
		budget *= 0.40
	} else if 10 <= people && people <= 24 {
		budget *= 0.50
	} else if 25 <= people && people <= 49 {
		budget *= 0.60
	} else {
		budget *= 0.75
	}

	ticketPrice := ticketsType[ticket] * people
	moneyLeft := budget - ticketPrice

	if moneyLeft > 0 {
		fmt.Printf("Yes! You have %.2f leva left.", moneyLeft)
	} else {
		fmt.Printf("Not enough money! You need %.2f leva.", math.Abs(moneyLeft))
	}

}

/*
name   :01. Match Tickets
input  :1000 Normal 1
output :Yes! You have 0.01 leva left.
*/
