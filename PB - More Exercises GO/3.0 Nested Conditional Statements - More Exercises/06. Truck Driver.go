package main

import "fmt"

func main() {
    var season, tariffType string
    var kmPerMonth float64
    fmt.Scan(&season, &kmPerMonth)

    seasonsPriceTable  := map[string]map[string]float64 {
        "minTariff" : {"Autumn" : 0.75, "Spring" : 0.75, "Summer" : 0.9, "Winter" : 1.05},
        "middleTariff" : {"Autumn" : 0.95,"Spring" : 0.95, "Summer" : 1.10, "Winter" : 1.25},
        "maxTariff" : {"Autumn" : 1.45,"Spring" : 1.45, "Summer" : 1.45, "Winter" : 1.45},
    }

    if kmPerMonth <= 5000 {
        tariffType = "minTariff"
    } else if 5000 < kmPerMonth && kmPerMonth <= 10000 {
        tariffType = "middleTariff"
    } else if 10000 < kmPerMonth && kmPerMonth <= 20000 {
        tariffType = "maxTariff"
    }

    countMount := 4.0
    tax := 0.90
    moneyPassedKm := (kmPerMonth * seasonsPriceTable[tariffType][season]) * countMount * tax

    fmt.Printf("%.2f",moneyPassedKm )
}

/*
name   :06. Truck Driver
input  :Summer 3455
output :11194.20
*/
