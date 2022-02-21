package main

import "fmt"

func main() {
    var vegetable, fruit, kgVegetable, kgFruit float64
    fmt.Scan(&vegetable, &fruit, &kgVegetable, &kgFruit)
    fmt.Printf("%.2f", (vegetable * kgVegetable + fruit * kgFruit) / 1.94)
}


/*
name   :04. Vegetable Market
input  :0.194 19.4 10 10
output :101.00
*/

