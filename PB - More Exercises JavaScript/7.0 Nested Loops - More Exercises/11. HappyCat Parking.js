function happyCatParking(arg) {
    let [days, hours] = [...arg.map(Number)]
    let [currentTax, totalSum] = [0, 0]

    for (let day = 1; day <= days; day++) {
        for (let hour = 1; hour <= hours; hour++) {
            if (!(day & 1) && hour & 1) {
                currentTax += 2.50
            } else if (day & 1 && !(hour & 1)) {
                currentTax += 1.25
            } else {
                currentTax += 1.00
            }

        }
        console.log(`Day: ${day} - ${currentTax.toFixed(2)} leva`)
        totalSum += currentTax
        currentTax = 0

    }

    console.log(`Total: ${totalSum.toFixed(2)} leva`)
}

// happyCatParking([2, 5])