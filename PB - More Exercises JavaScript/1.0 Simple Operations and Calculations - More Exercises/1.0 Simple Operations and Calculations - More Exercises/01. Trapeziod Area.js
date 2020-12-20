function trapeziodArea(...arg) {
    let [b1, b2, h,] = [...arg.map(Number)]
    return ((b1 + b2) * h / 2).toFixed(2)
}

//console.log(trapeziodArea(8, 13, 7)) //73.50
