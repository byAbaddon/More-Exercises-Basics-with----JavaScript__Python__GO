function backToThePast(arg) {
    let [money, years] = [...arg.map(Number)]
    let years_old = 17

    for (let i = 1800; i <= years; i++) {
        years_old += 1

        if (i & 1) {
            money -= 12000 + years_old * 50
        } else {
            money -= 12000
        }
    }

    if (money >= 0)
        return `Yes! He will live a carefree life and will have ${money.toFixed(2)} dollars left.`
    return `He will need ${Math.abs(money).toFixed(2)} dollars to survive.`
}

//console.log(backToThePast([50000, 1802]))  //Yes! He will live a carefree life and will have 13050.00 dollars left.
