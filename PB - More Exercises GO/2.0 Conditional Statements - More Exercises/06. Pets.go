package main

import "fmt"
import "math"

func main() {
	var days, kgFood, dogFoodKg, catFoodKg, turtleFoodKg float64
	fmt.Scan(&days, &kgFood, &dogFoodKg, &catFoodKg, &turtleFoodKg)
	food := (days * dogFoodKg) + (days * catFoodKg) + (days*turtleFoodKg)/1000
	if food <= kgFood {
		fmt.Printf("%v kilos of food left.", math.Floor(kgFood-food))
	} else {
		fmt.Printf("%v more kilos of food are needed.", math.Ceil(food-kgFood))

	}
}

/*
name   :06. Pets

input  :2 10 1 1 1200
output :3 kilos of food left.

input  :5 10 2.1 0.8 321
output :7 more kilos of food are needed.
*/
