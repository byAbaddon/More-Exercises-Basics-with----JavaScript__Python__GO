function vacation(budget, season) {
    let [location, place_type] = ['', '']

    if (budget <= 1000) {
        place_type = 'Camp'
        if (season == 'Summer') {
            location = 'Alaska'
            budget *= 0.65
        } else {
            location = 'Morocco'
            budget *= 0.45
        }
    } else if (1000 < budget && budget <= 3000) {
        place_type = 'Hut'
        if (season == 'Summer') {
            location = 'Alaska'
            budget *= 0.80
        } else {
            location = 'Morocco'
            budget *= 0.60
        }
    } else {
        place_type = 'Hotel'
        if (season == 'Summer') {
            location = 'Alaska'
            budget *= 0.90
        } else {
            location = 'Morocco'
            budget *= 0.90
        }
    }

    return `${location} - ${place_type} - ${budget.toFixed(2)}`
}

// console.log(vacation(800, 'Summer'))   //Alaska - Camp - 520.00
