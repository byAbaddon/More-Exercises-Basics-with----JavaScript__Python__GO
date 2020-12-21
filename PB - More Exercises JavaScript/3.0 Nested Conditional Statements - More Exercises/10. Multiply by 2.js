function multiplyby2(...arg) {

    while (arg) {
        let currentNum = arg.shift()
        if (currentNum < 0) {
            console.log('Negative number!')
            break
        }
        console.log(`Result: ${(currentNum * 2).toFixed(2)}`)  
    }
}

// multiplyby2(12, 43,2144, 12.3, 543.23, -20)