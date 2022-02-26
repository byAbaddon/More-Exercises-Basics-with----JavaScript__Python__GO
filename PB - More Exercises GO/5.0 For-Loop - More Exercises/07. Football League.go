package main

import "fmt"

func main() {
    var stadium, fens, A, B, V, G float64
    fmt.Scan(&stadium, &fens)
    
    for i := 0; i < int(fens); i++ {
        var sector string
        fmt.Scan(&sector)
        
        switch sector {
        case "A": A++
        case "B": B++
        case "V": V++
        case "G": G++            
        }   
    }
    
    fmt.Printf("%.2f%%\n",A / fens * 100)
    fmt.Printf("%.2f%%\n",B / fens * 100)
    fmt.Printf("%.2f%%\n",V / fens * 100)
    fmt.Printf("%.2f%%\n",G / fens * 100)
    fmt.Printf("%.2f%%" ,fens / stadium * 100)
}

/*
name   :07. Football League
input  : 76 10 A V V V G B A V B B
output :
20.00%
30.00%
40.00%
10.00%
13.16%
*/
