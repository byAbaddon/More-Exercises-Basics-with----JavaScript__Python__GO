package main

import "fmt"

func main() {
	var skumriqPrice, cacaPrice, kgPalamud, kgSafrid, kgMidi float64
	fmt.Scan(&skumriqPrice, &cacaPrice, &kgPalamud, &kgSafrid, &kgMidi)

	palamudPrice := skumriqPrice + 0.6*skumriqPrice
	safridPrice := cacaPrice + 0.8*cacaPrice
	price := (kgPalamud*palamudPrice + kgSafrid*safridPrice + kgMidi*7.5)
	fmt.Printf("%.2f", price)
}

/*
name   :06. Fishland
input  :6.90 4.20 1.5 2.5 1
output :42.96
*/
