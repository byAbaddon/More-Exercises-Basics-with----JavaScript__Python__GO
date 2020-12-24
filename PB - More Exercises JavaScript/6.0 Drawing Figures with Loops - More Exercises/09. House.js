function house(n) {
    n = +n
    if (n === 2) return `**\n||`

    //top
    if (!(n & 1)) {
        for (i = 1; i < (n / 2); i++) { //is even
            console.log('-'.repeat(n / 2 - i) + '*'.repeat(i * 2) + '-'.repeat((n / 2 - i)))
        }
    } else {
        for (i = 1; i < (n / 2); i++) { //is odd
            console.log('-'.repeat(Math.ceil(n / 2) - i) + '*'.repeat(i * 2 - 1) + '-'.repeat(Math.ceil(n / 2) - i))
        }
    }

    //middle part
    console.log('*'.repeat(n))

    //bottom
    if (!(n & 1)) {
        for (i = 0; i < n / 2; i++) {
            console.log('|' + '*'.repeat(n - 2) + '|')
        }
    }

    if (n & 1) {
        for (i = 0; i < n / 2 - 1; i++) {
            console.log('|' + '*'.repeat(n - 2) + '|');
        }
    }
}

// house(4)
// console.log(house(2))
