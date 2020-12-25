function carNumbers(arg) {
    let [start, end] = [...arg.map(Number)]
    let result = ''

    for (let a = start; a <= end; a++) {
        for (let b = start; b <= end; b++) {
            for (let c = start; c <= end; c++) {
                for (let d = start; d <= end; d++) {
                    if (!(a & 1) && d & 1 || a & 1 && !(d & 1)) {
                        if (a > d) {
                            if (!((b + c) & 1)) {
                                result += '' + a + b + c + d + ' '
                            }
                        }
                    }
                }
            }
        }
    }

    return result
}

//console.log(carNumbers([2, 3]))

