function pets(...arg) {
    let [days, kgFood, dogFoodKg, catFoodKg, turtleFoodKg] = [...arg.map(Number)]
    let food = (days * dogFoodKg) + (days * catFoodKg) + (days * turtleFoodKg) / 1000

    if (food <= kgFood)
        return `${Math.floor(kgFood - food)} kilos of food left.`
    return `${Math.ceil(food - kgFood)} more kilos of food are needed.`
}

//console.log(pets(5, 10, 2.1, 0.8, 321))    //7 more kilos of food are needed.
