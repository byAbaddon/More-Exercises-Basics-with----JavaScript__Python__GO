function challengeTheWedding(arg) {
    let [men, women, table] = [...arg.map(Number)]
    const visitedList = []
    for (let m = 1; m <= men; m++) {
        for (let w = 1; w <= women; w++) {
            if (table !== visitedList.length) {
                visitedList.push(`(${m} <-> ${w})`)
            } else {
                break
            }
        }
    }

    return visitedList.join(' ')
}

//console.log(challengeTheWedding([2, 2, 6]))