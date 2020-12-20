function fuelTankTwo(typeFuel, quantity, clubCard) {
    const fuel_price_dict = {'Gasoline': 2.22,'Diesel': 2.33,'Gas': 0.93}
    const discount = {'Gasoline': 0.18,'Diesel': 0.12,'Gas': 0.08}
    let subtotalPrice = 0
    let totalPrice = 0

    if (clubCard === 'Yes') {
        subtotalPrice = fuel_price_dict[typeFuel] - discount[typeFuel]
        totalPrice = subtotalPrice * quantity
    } else {
        totalPrice = fuel_price_dict[typeFuel] * quantity
    }

    if (20 <= quantity && quantity <= 25) {
        totalPrice *= 0.92
    } else if (quantity > 25) {
        totalPrice *= 0.90
    }

    return totalPrice.toFixed(2) + ' lv.'
}

// console.log(fuelTankTwo('Gas', 30, 'Yes'))  //22.95 lv.
