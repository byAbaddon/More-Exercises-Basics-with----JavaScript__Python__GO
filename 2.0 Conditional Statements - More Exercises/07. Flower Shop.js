function flowerShop(...arg) {
    let [countMgn, countZum, countRoses, countCactus, giftPrice] = [...arg.map(Number)]
    let tax = 0.95
    let totalSum = (countMgn * 3.25 + countZum * 4 + countRoses * 3.5 + countCactus * 8) * tax

    if (totalSum < giftPrice)
        return `She will have to borrow ${Math.ceil(giftPrice - totalSum)} leva.`
    return `She is left with ${~~(totalSum - giftPrice)} leva.`
}

// console.log(flowerShop(2, 3, 5, 1, 50)) //She will have to borrow 9 leva.