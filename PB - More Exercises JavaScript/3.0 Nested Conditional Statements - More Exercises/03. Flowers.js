function flowers(...arg) {
    let [chrysanthemums, roses, tulips, season, isCelebrateDay] = [...arg.map(el => isNaN(el) ? el : Number(el))]
    let [chrysanthemumTotal, roseTotal, tulipTotal] = [0, 0, 0]

    if (season === 'Spring' || season === 'Summer') {
        chrysanthemumTotal = chrysanthemums * 2
        roseTotal = roses * 4.1
        tulipTotal = tulips * 2.5
    } else if (season === 'Autumn' || season === 'Winter') {
        chrysanthemumTotal = chrysanthemums * 3.75
        roseTotal = roses * 4.5
        tulipTotal = tulips * 4.15
    }

    let price = chrysanthemumTotal + roseTotal + tulipTotal

    if (isCelebrateDay == 'Y')
        price = price * 1.15

    if (season === 'Spring' && tulips > 7)
        price *= 0.95

    if (season == 'Winter' && roses >= 10)
        price *= 0.90

    if (chrysanthemums + roses + tulips > 20)
        price *= 0.8

    return (price + 2).toFixed(2)
}

// console.log(flowers(2, 4, 8, 'Spring', 'Y', ))   //46.14