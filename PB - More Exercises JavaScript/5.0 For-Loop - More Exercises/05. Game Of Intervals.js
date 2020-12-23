function gameOfIntervals(arg) {
    arg = [...arg.map(Number)]
    let move = arg.shift()
    let points = 0
    const percentObj = {'From 0 to 9:': 0,'From 10 to 19:': 0,'From 20 to 29:': 0,'From 30 to 39:': 0,'From 40 to 50:': 0,'Invalid numbers:': 0}

    for (let i = 0; i < move; i++) {
        let currentNum = arg[i]
        if (0 <= currentNum && currentNum <= 9) {
            points += currentNum * 0.20
            percentObj['From 0 to 9:'] += 1
        } else if (10 <= currentNum && currentNum <= 19) {
            points += currentNum * 0.30
            percentObj['From 10 to 19:'] += 1
        } else if (20 <= currentNum && currentNum <= 29) {
            points += currentNum * 0.40
            percentObj['From 20 to 29:'] += 1
        } else if (30 <= currentNum && currentNum <= 39) {
            points += 50
            percentObj['From 30 to 39:'] += 1
        } else if (40 <= currentNum && currentNum <= 50) {
            points += 100
            percentObj['From 40 to 50:'] += 1
        } else {
            points /= 2
            percentObj['Invalid numbers:'] += 1
        }
    }

    console.log(points.toFixed(2))

    for (const [key, val] of Object.entries(percentObj)) {
        console.log(`${key} ${(val / move * 100).toFixed(2)}%`)
    }
}

// gameOfIntervals([10, 43, 57, -12, 23, 12, 0, 50, 40, 30, 20])