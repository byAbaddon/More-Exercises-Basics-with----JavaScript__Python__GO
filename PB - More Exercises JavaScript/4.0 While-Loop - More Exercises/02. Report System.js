function reportSystem([...arg]) {
    if (isNaN(arg[arg.length - 1])) arg.pop()

    let necessaryMoney = +arg.shift()
    let cashPaymentCount = 0
    let cashPayment = 0
    let cardPayment = 0
    let cardPaymentCount = 0
    let totalMoney = 0
    let isPaymentCash = true

    for (let i = 1; i <= arg.length; i++) {
        let products_price = +arg[i - 1]

        if (isPaymentCash) {
            isPaymentCash = false
            if (products_price > 100) {
                console.log('Error in transaction!')
            } else {
                cashPayment += products_price
                cashPaymentCount += 1
                console.log('Product sold!')
                totalMoney += products_price
            }
        } else {
            isPaymentCash = true
            if (products_price < 10) {
                console.log('Error in transaction!')
            } else {
                cardPayment += products_price
                cardPaymentCount += 1
                console.log('Product sold!')
                totalMoney += products_price
            }

            if (totalMoney >= necessaryMoney)
                break
        }
    }

    if (totalMoney < necessaryMoney) {
        console.log('Failed to collect required money for charity.')
    } else {
        let averageCashPayment = cashPayment / cashPaymentCount
        let averageCardPayment = cardPayment / cardPaymentCount
        console.log(`Average CS: ${averageCashPayment.toFixed(2)}`)
        console.log(`Average CC: ${averageCardPayment.toFixed(2)}`)
    }
}


// reportSystem([500, 120, 8, 63, 256, 78, 317])
// reportSystem([600, 86, 150, 98, 227, 'End'])