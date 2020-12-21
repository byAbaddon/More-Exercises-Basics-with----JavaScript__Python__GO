function truckDriver(season, kmPerMonth) {
    let tariffType = ''
    const seasons_price_table_dict = {
        'min_tariff' : {'Autumn' : 0.75, 'Spring' : 0.75, 'Summer' : 0.9, 'Winter' : 1.05},
        'middle_tariff' : {'Autumn' : 0.95,'Spring' : 0.95, 'Summer' : 1.10, 'Winter' : 1.25},
        'max_tariff' : {'Autumn' : 1.45,'Spring' : 1.45, 'Summer' : 1.45, 'Winter' : 1.45},
    }

    if (kmPerMonth <= 5000) {
        tariffType = 'min_tariff'
    } else if (5000 < kmPerMonth && kmPerMonth <= 10000) {
        tariffType = 'middle_tariff'
    } else if (10000 < kmPerMonth <= 20000) {
        tariffType = 'max_tariff'
    }

    let countMount = 4
    let tax = 0.90
    let moneyPassedKm = (kmPerMonth * seasons_price_table_dict[tariffType][season]) * countMount * tax

    return `${moneyPassedKm.toFixed(2)}`
}

// console.log(truckDriver('Winter', 16042))    //83739.24
