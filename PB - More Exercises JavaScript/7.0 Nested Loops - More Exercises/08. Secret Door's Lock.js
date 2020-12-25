function secretDoorsLock(arg) {
    let [hundredNum, doubleNum, singleNum] = [...arg.map(Number)]
    const primeList = [2, 3, 5, 7]

    for (let h = 2; h <= hundredNum; h += 2) {
        for (let d = 2; d <= doubleNum; d++) {
            for (let s = 2; s <= singleNum; s += 2) {
                if (primeList.includes(d)) {
                    console.log(`${h} ${d} ${s}`);
                }
            }
        }
    }
}

//secretDoorsLock([2, 5, 5])
