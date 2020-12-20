function pipesInPool(...arg) {
    let [volume, p1, p2, hours] = [...arg.map(Number)]
    let pool = (p1 + p2) * hours
    let poolPercent = (pool / volume) * 100
    let p1Percent = (p1 * hours / pool) * 100
    let p2Percent = (p2 * hours / pool) * 100
    let overFlow = pool - volume

    if (pool > volume)
        return `For ${hours} hours the pool overflows with ${parseFloat(overFlow.toFixed(2) )} liters.`
    return `The pool is ${poolPercent.toFixed(2)}% full. Pipe 1: ${p1Percent.toFixed(2  )}%. Pipe 2: ${p2Percent.toFixed(2)}%.`
}

// console.log(pipesInPool(1000, 100, 120, 3))