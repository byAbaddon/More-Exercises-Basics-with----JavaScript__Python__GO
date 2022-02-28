package main

import "fmt"

func main() {
	var password, result string
	var pass, passCount int
	fmt.Scan(&pass)

	for a := 1; a <= 9; a++ {
		for b := 1; b <= 9; b++ {
			for c := 1; c <= 9; c++ {
				for d := 1; d <= 9; d++ {
					if pass == a*b+c*d && (a < b && c > d) {
						result += fmt.Sprintf("%v%v%v%v ", a, b, c, d)
						passCount++
						if passCount == 4 {
							password = fmt.Sprintf("%v%v%v%v", a, b, c, d)
						}
					}
				}
			}
		}
	}

	fmt.Println(result)

	if len(password) > 1 {
		fmt.Println("Password:", password)
	} else {
		fmt.Println("No!")
	}

}

/*
name   :12. The song of the wheels
input  :11
output :
1291 1342 1381 1471 1532 1561 1651 1741 1831 1921 2351 2431
Password: 1471
*/
