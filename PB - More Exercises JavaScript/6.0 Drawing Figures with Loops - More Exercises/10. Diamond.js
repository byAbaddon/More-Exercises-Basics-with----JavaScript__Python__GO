function diamond(n) {
    n = +n
    if (n === 1) return '*'
    if (n === 2) return '**'

    let stars = 1
    if (!(n & 1)) {
        stars = 2
    }

    // up
    for (i = 1; i <= 1; i++) {
        console.log('-'.repeat((Math.ceil(n / 2 - 1))) + '*'.repeat(stars) + '-'.repeat(Math.ceil(n / 2 - 1)))
    }

    //up  /odd
    if (n > 4 && n % 2 !== 0) {
        for (i = 1; i <= (n / 2 - 1); i++) {
            console.log('-'.repeat((n / 2 - i)) + '*' + '-'.repeat((i + i) - 1) + '*' + '-'.repeat(((n / 2 - i))))
        }
    } else {
        for (i = 1; i <= (n / 2 - 2); i++) { //even
            console.log('-'.repeat(((n / 2 - i) - 1)) + '*' + '-'.repeat((i + i)) + '*' + '-'.repeat(((n / 2 - i) - 1)))
        }
    }

    // middle line
    for (i = 0; i < 1; i++) {
        console.log('*' + '-'.repeat((n - 2)) + '*')
    }

    //under middle
    if (n > 4 && n % 2 !== 0) {
        for (i = (n / 2 - 1); i >= 1; i--) {
            console.log('-'.repeat((n / 2 - i)) + '*' + '-'.repeat((i + i) - 2) + '*' + '-'.repeat(((n / 2 - i))))
        }
    } else {
        for (i = (n / 2 - 2); i >= 1; i--) { //even
            console.log('-'.repeat(((n / 2 - i) - 1)) + '*' + '-'.repeat((i + i)) + '*' + '-'.repeat(((n / 2 - i) - 1)))
        }
    }

    // bottom
    for (i = 1; i <= 1; i++) {
        console.log('-'.repeat((Math.ceil(n / 2 - 1))) + '*'.repeat(stars) + '-'.repeat(Math.ceil(n / 2 - 1)))
    }
}

// diamond(3)
// diamond(5)