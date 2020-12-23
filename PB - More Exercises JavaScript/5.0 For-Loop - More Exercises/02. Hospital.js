function hospital(arg) {
    arg = [...arg.map(Number)]
    let period = arg.shift()
    let = [untreatedPatients, treatedPatients, doctors] = [0, 0, 7]

    for (let i = 1; i <= period; i++) {
        let patientList = arg[i - 1]

        if (i % 3 == 0) {
            if (untreatedPatients > treatedPatients) {
                doctors += 1
            }
        }

        if (0 <= patientList && patientList <= doctors) {
            treatedPatients += patientList
        } else {
            untreatedPatients = (patientList - doctors) + untreatedPatients
            treatedPatients += doctors
        }
    }

    return `Treated patients: ${treatedPatients}.\nUntreated patients: ${untreatedPatients}.`
}

//console.log(hospital([4, 7, 27, 9, 1]))    //Treated patients: 23. , Untreated patients: 21.

