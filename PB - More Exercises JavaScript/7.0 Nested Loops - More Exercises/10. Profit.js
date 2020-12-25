function profit(arg) {
    let [oneCoins, twoCoins, fiveCoins, allSum] = [...arg.map(Number)]

    for (let one = 0; one <= oneCoins; one++) {
        for (let two = 0; two <= twoCoins; two++) {
            for (let five = 0; five <= fiveCoins; five++) {
                if (one * 1 + two * 2 + five * 5 === allSum)
                    console.log(`${one} * 1 lv. + ${two} * 2 lv. + ${five} * 5 lv. = ${allSum} lv.`)
            }
        }
    }
}

//profit([3, 2, 3, 10])

