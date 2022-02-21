package main

import "fmt"

func main() {
    var x, y, h float64
    fmt.Scan(&x, &y, &h)
    doorArea := 2 * 1.2
    windowsArea := 2 * (1.5 * 1.5)
    frontWallsArea := 2 * (x * x) - doorArea
    sideWalls := 2 * (x * y) - windowsArea
    totalWallArea := frontWallsArea + sideWalls
    greenPaint := totalWallArea / 3.4
    roof_square_area := 2 * (x * y)
    roofTriangleArea := (x * h)
    totalRoofArea := roof_square_area + roofTriangleArea
    redPaint := totalRoofArea / 4.3
    
    fmt.Printf("%.2f\n%.2f",greenPaint, redPaint )
}

/*
name   :07. House Painting
input  :6 10 5.2
output :54.44 / 35.16
*/
