function vegetableMarket(...arg) {
    let [vegetable, fruit, kgVegetable, kgFruit] = [...arg.map(Number)]
    return ((vegetable * kgVegetable + fruit * kgFruit) / 1.94).toFixed(2)
}

// console.log(vegetableMarket(0.194, 19.4, 10, 10))  //101.00



