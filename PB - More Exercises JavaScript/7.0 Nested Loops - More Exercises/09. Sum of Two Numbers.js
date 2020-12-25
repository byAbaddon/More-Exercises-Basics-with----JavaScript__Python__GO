function sumOfTwoNumbers(arg) {
    let [start, end, magicNum] = [...arg.map(Number)]
    let counter = 0
    
    for (let a = start; a <= end; a++) {
        for (let b = start; b <= end; b++) {
            counter++
            if (a + b === magicNum) {
                return `Combination N:${counter} (${a} + ${b} = ${a + b})`
            }
        }
    }

    return `${counter} combinations - neither equals ${magicNum}`
}

// console.log(sumOfTwoNumbers([1, 10, 5]))
// console.log(sumOfTwoNumbers([23, 24, 20]))