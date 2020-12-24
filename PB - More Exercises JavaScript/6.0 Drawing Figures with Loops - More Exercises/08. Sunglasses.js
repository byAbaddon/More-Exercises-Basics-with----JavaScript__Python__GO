function sunglasses(n) {
    n = +n
    //top
    console.log(`${'*'.repeat(2 * n)}${' '.repeat(n) }${'*'.repeat(2 * n)}`)

    //middle
    for (i = 0; i < n - 2; i++) {
        if (i === ~~((n - 1) / 2 - 1)) {
            console.log('*' + '/'.repeat(n + (n - 2)) + '*' + '|'.repeat(n) + '*' + '/'.repeat(n + (n - 2)) + '*')
        } else {
            console.log('*' + '/'.repeat(n + (n - 2)) + '*' + ' '.repeat(n) + '*' + '/'.repeat(n + (n - 2)) + '*')
        }
    }

    //bottom
    console.log(`${'*'.repeat(2 * n)}${' '.repeat(n) }${'*'.repeat(2 * n)}`)
}

 //sunglasses(3)
// sunglasses(5)