function lettersCombinations(arg) {
    let [lOne, lTwo, lIgnore] = [...arg.map((el, index) => index < 2 ? el.charCodeAt(0) : el)]
    let collectionList = []

    for (let a = lOne; a <= lTwo; a++) {
        for (let b = lOne; b <= lTwo; b++) {
            for (let c = lOne; c <= lTwo; c++) {

                let current = String.fromCharCode(a) + String.fromCharCode(b) + String.fromCharCode(c)
                if (!current.includes(lIgnore)) {
                    collectionList.push(current)
                }
            }
        }
    }

    return `${collectionList.join(' ')} ${collectionList.length}`
}

// console.log(lettersCombinations(['a', 'c', 'b']))

