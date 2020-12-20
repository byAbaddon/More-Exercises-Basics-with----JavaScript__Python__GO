function fuelTank(typeFuel, quantity) {
    const lst = ['Diesel', 'Gasoline', 'Gas']

    if (!lst.includes(typeFuel))
        return 'Invalid fuel!'
    if (quantity >= 25)
        return `You have enough ${typeFuel.toLowerCase()}.`
    if (quantity < 25)
        return `Fill your tank with ${typeFuel.toLowerCase()}!`
}

// console.log(fuelTank('Diesel', 10))   //Fill your tank with diesel!
