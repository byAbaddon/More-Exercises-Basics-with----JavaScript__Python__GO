function christmasTree(n) {
    //head
    for (let i = 0; i < 1; i++) {
        console.log(`${' '.repeat(n)} | ${' '.repeat(n)}`)
    }

    //body
    for (let i = 1; i <= n; i++) {
        console.log(`${' '.repeat(n - i)}${'*'.repeat(i)} | ${'*'.repeat(i)}`)
    }
}

// christmasTree(2)
