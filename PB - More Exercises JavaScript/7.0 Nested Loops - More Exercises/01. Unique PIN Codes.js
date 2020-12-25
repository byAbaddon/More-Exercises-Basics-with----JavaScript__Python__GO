function uniquePINCodes(arg) {
    let [oneMax, twoMax, threeMax] = [...arg.map(Number)]
    const primeNumberList = [2, 3, 5, 7]
    for (let f = 2; f <= oneMax; f += 2) {
        for (let s = 2; s <= twoMax; s++) {
            for (let t = 2; t <= threeMax; t += 2) {
                if (primeNumberList.includes(s)) {
                    console.log(f, s, t);
                }
            }
        }
    }
}

//uniquePINCodes([3, 5, 5])
