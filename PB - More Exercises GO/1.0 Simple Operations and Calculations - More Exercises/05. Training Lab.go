package main

import "fmt"

func main() {
	var height, width float32
	fmt.Scan(&height, &width)
	hallway := 100
	lostWorkPosition := 3.0
	workPositionWidth := 70.0
	workPositionHeight := 120

	height *= 100
	width *= 100
	width -= float32(hallway)

	rest := int(width) % int(workPositionWidth)
	width -= float32(rest)
	bureauCount := width / float32(workPositionWidth)

	rest = int(height) % workPositionHeight
	height -= float32(rest)
	rowCount := height / float32(workPositionHeight)

	fmt.Println(bureauCount*rowCount - float32(lostWorkPosition))
}

/*
name   :05. Training Lab
input  :15 8.9
output :129
*/

