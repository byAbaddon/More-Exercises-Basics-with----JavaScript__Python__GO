function carToGo(budget, season) {
    let [classType, carType] = ['', '']

    if (budget <= 100) {
        classType = 'Economy class'
        if (season == 'Summer') {
            budget *= 0.35
            carType = 'Cabrio'
        } else {
            budget *= 0.65
            carType = 'Jeep'
        }
    } else if (100 < budget && budget <= 500) {
        classType = 'Compact class'
        if (season == 'Summer') {
            budget *= 0.45
            carType = 'Cabrio'
        } else {
            budget *= 0.80
            carType = 'Jeep'
        }
    } else {
        classType = 'Luxury class'
        budget *= 0.90
        carType = 'Jeep'
    }

    return `${classType}\n${carType} - ${budget.toFixed(2)}`
}

// console.log(carToGo(450, 'Summer'))       //Compact class ; Cabrio - 202.50

