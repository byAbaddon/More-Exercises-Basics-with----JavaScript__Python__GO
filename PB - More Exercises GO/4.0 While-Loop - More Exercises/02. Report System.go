package main

import "fmt"
import "strconv"

func main() {
	var necessaryMoney float64
	fmt.Scan(&necessaryMoney)
	var cashPaymentCount, cashPayment, cardPayment, cardPaymentCount, totalMoney float64
	var isPaymentCash = true

	for {
		var token string
		fmt.Scan(&token)

		if token == "End" {
			break
		}

		products_price, _ := strconv.ParseFloat(token, 8)

		if isPaymentCash {
			isPaymentCash = false
			if products_price > 100 {
				fmt.Println("Error in transaction!")
			} else {
				cashPayment += products_price
				cashPaymentCount += 1
				fmt.Println("Product sold!")
				totalMoney += products_price
			}
		} else {
			isPaymentCash = true
			if products_price < 10 {
				fmt.Println("Error in transaction!")
			} else {
				cardPayment += products_price
				cardPaymentCount += 1
				fmt.Println("Product sold!")
				totalMoney += products_price
			}

			if totalMoney >= necessaryMoney {
				break
			}
		}
	}

	if totalMoney < necessaryMoney {
		fmt.Println("Failed to collect required money for charity.")
	} else {
		csAverageCashPayment := cashPayment / cashPaymentCount
        ccAverageCardPayment := cardPayment / cardPaymentCount
		fmt.Printf("Average CS: %.2f\n", csAverageCashPayment)
		fmt.Printf("Average CC: %.2f", ccAverageCardPayment)
	}

}

/*
name   :02. Report System
input  :500 120 8 63 256 78 317
output :
Error in transaction!
Error in transaction!
Product sold!
Product sold!
Product sold!
Product sold!
Average CS: 70.50
Average CC: 286.50
*/
