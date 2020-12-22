function dishwasher([...arg]) {
    if (isNaN(arg[arg.length-1])) arg.pop()

    let detergent = +arg.shift() * 750
    let [dishes, pots] = [0, 0]

    for (let i = 1; i <= arg.length; i++) {
        let current = +arg[i - 1]
        if (i % 3 != 0) {
            detergent -= current * 5
            dishes += current
        } else {
            detergent -= current * 15
            pots += current
        }
    }

    if (detergent >= 0)
        return `Detergent was enough!\n${dishes} dishes and ${pots} pots were washed.\nLeftover detergent ${detergent} ml.`
    return `Not enough detergent, ${Math.abs(detergent)} ml. more necessary!`
}

//console.log(dishwasher([2, 53, 65, 55, 'End']))
//console.log(dishwasher([1, 10, 15, 10, 12, 13, 30]))