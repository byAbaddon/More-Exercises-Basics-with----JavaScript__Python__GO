package main

import "fmt"

func main() {
	var cargoCount int
	fmt.Scan(&cargoCount)
	cargosObj := map[string]float64{"allCargo": 0.0, "bus": 0.0, "truck": 0.0, "train": 0.0}

	for i := 0; i < cargoCount; i++ {
		var currentCargo float64
		fmt.Scan(&currentCargo)

		if currentCargo <= 3 {
			cargosObj["bus"] += currentCargo
		} else if 4 <= currentCargo && currentCargo <= 11 {
			cargosObj["truck"] += currentCargo
		} else {
			cargosObj["train"] += currentCargo
		}

		cargosObj["allCargo"] += currentCargo
	}

	averagePrice := (cargosObj["bus"]*200 + cargosObj["truck"]*175 + cargosObj["train"]*120) / cargosObj["allCargo"]

	fmt.Printf("%.2f\n", float64(averagePrice))
	fmt.Printf("%.2f%%\n", cargosObj["bus"]/cargosObj["allCargo"]*100)
	fmt.Printf("%.2f%%\n", cargosObj["truck"]/cargosObj["allCargo"]*100)
	fmt.Printf("%.2f%%\n", cargosObj["train"]/cargosObj["allCargo"]*100)
}

/*
name   :03. Logistics
input  :4 1 5 16 3
output :
143.80
16.00%
20.00%
64.00%
*/
