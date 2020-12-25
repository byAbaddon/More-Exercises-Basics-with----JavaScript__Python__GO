function theSongOfTheWheels(pass) {
    let [passCount, password, result] = [0, '', '']

    for (let a = 1; a <= 9; a++) {
        for (let b = 1; b <= 9; b++) {
            for (let c = 1; c <= 9; c++) {
                for (let d = 1; d <= 9; d++) {
                    if (+pass === a * b + c * d && (a < b && c > d)) {
                        result += `${a}${b}${c}${d} `
                        passCount++
                        if (passCount === 4) {
                            password = `${a}${b}${c}${d}`
                        }
                    }
                }
            }
        }
    }

    console.log(result.trim())

    if (password)
        return `Password: ${password}`
    return 'No!'
}


// console.log(theSongOfTheWheels([11]))
// console.log(theSongOfTheWheels([139]))