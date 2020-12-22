function averageNumber([...arg]) {
    let sum = 0
    let counter = 0
    for (let i = 1; i < arg.length; i++) {
        sum += +arg[i]
        counter++
    }

    return (sum / counter).toFixed(2) 
}

// console.log(averageNumber([2, 4, 6]))
// console.log(averageNumber([4, 3, 2, 4, 2]))