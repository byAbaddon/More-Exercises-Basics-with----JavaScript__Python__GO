function safePasswordsGenerator(arg) {
    let [numA, numB, max] = [...arg.map(Number)]
    let [codeA, codeB, password] = [35, 64, '']

    while (codeA <= 55 && codeB <= 96) {
        for (let x = 1; x <= numA; x++) {
            for (let y = 1; y <= numB; y++) {
                let charA = String.fromCharCode(codeA)
                let charB = String.fromCharCode(codeB)

                password += `${charA}${charB}${x}${y}${charB}${charA}|`
                max--
                codeA++
                codeB++
                codeA > 55 ? codeA = 35 : null
                codeB > 96 ? codeB = 64 : null

                if (max === 0 || (x === numA && y === numB)) {
                    return password
                    
                }
            }
        }
    }
}

// safePasswordsGenerator(['2', '3', '10'])