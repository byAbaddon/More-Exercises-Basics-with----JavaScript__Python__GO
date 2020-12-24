function rhombusOfStars(n) {
    //first part
    for (let i = 1; i <= n; i++) {
        console.log(`${' '.repeat(n - i) }${'* '.repeat(i)}`)
    }

    //second part
    for (let i = 1; i < n; i++) {
        console.log(`${' '.repeat(i) }${'* '.repeat(n -i)}`)
    }
}

// rhombusOfStars(4)