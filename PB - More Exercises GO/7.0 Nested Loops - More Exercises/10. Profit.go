package main

import "fmt"

func main() {
	var oneCoins, twoCoins, fiveCoins, allSum int
	fmt.Scan(&oneCoins, &twoCoins, &fiveCoins, &allSum)

	for one := 0; one <= oneCoins; one++ {
		for two := 0; two <= twoCoins; two++ {
			for five := 0; five <= fiveCoins; five++ {
				if one*1+two*2+five*5 == allSum {
					fmt.Printf("%v * 1 lv. + %v * 2 lv. + %v * 5 lv. = %v lv.\n", one, two, five, allSum)
				}
			}
		}
	}

}

/*
name   :10. Profit
input  :3 2 3 10
output :
0 * 1 lv. + 0 * 2 lv. + 2 * 5 lv. = 10 lv.
1 * 1 lv. + 2 * 2 lv. + 1 * 5 lv. = 10 lv.
3 * 1 lv. + 1 * 2 lv. + 1 * 5 lv. = 10 lv.
*/
