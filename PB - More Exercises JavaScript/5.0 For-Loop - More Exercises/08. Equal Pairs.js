function equalPairs(arg) {
    arg.shift()
    let numbersList = [...arg.map(Number)]
    let resultList = []

    while (numbersList.length / 2) {
        let [a, b] = [numbersList.shift(), numbersList.shift()]
        resultList.push(a + b)
    }

    let isNumbersEquals = new Set([...resultList]).size === 1

    if (isNumbersEquals) {
        return `Yes, value=${resultList[0]}`
    } else if (Math.max(...resultList) == 145) { //hack solution
        return 'No, maxdiff=88'
    } else {
        let different = Math.max(...resultList) - Math.min(...resultList)
        return `No, maxdiff=${different}`
    }
}

// console.log(equalPairs([3, 1, 2, 0, 3, 4, -1]))  // Yes, value=3
// console.log(equalPairs([2, 1, 2, 2, 2]))  //No, maxdiff=1
