function bills(arg) {
    arg = [...arg.map(Number)]
    let monthsCount = arg.shift()
    let allElectricity = arg.reduce((a, b) => a + b)

    let water = 20 * monthsCount
    let internet = 15 * monthsCount
    let other = allElectricity * 1.20 + water * 1.20 + internet * 1.2
    let average = (allElectricity + water + internet + other) / monthsCount

    console.log(`Electricity: ${allElectricity.toFixed(2)} lv`)
    console.log(`Water: ${water.toFixed(2)} lv`)
    console.log(`Internet: ${internet.toFixed(2)} lv`)
    console.log(`Other: ${other.toFixed(2)} lv`)
    console.log(`Average: ${average.toFixed(2)} lv`)
}

// bills([5, 68.63, 89.25, 132.53, 93.53, 63.22])