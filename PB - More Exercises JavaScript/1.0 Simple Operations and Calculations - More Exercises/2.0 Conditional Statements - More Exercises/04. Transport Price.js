function transportPrice(km, day) {
    price = 0

    if (km < 20) {
        if (day == 'day') {
            price = 0.79 * km + 0.70
        } else {
            price = 0.90 * km + 0.70
        }
    } else if (20 <= km  && km < 100) {
        price = km * 0.09
    } else {
        price = km * 0.06
    }

    return price.toFixed(2)
}

// console.log(transportPrice(7, 'night'))   //7.00