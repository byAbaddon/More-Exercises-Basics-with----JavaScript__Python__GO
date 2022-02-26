package main

import "fmt"

func main() {
	var period int
	fmt.Scan(&period)
	var untreatedPatients, treatedPatients, doctors = 0, 0, 7

	for i := 1; i <= period; i++ {
		var patientList int
		fmt.Scan(&patientList)

		if i%3 == 0 {
			if untreatedPatients > treatedPatients {
				doctors += 1
			}
		}

		if 0 <= patientList && patientList <= doctors {
			treatedPatients += patientList
		} else {
			untreatedPatients = (patientList - doctors) + untreatedPatients
			treatedPatients += doctors
		}
	}

	fmt.Printf("Treated patients: %v.\nUntreated patients: %v.", treatedPatients, untreatedPatients)
}

/*
name   :02. Hospital
input  :4 7 27 9 1
output :Treated patients: 23. , Untreated patients: 21.
*/
