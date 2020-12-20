function firm(...arg) {
    let [hours, day, workers] = [...arg]
    let remnantTime = ~~(day * 0.90 * 8)
    let extraWorkTime = workers * (2 * day)
    let totalTime = remnantTime + extraWorkTime

    if (totalTime >= hours)
        return `Yes!${totalTime - hours} hours left.`
    return `Not enough time!${hours - totalTime} hours needed.`
}

//console.log(firm(90, 7, 3))      //Yes!2 hours left.
