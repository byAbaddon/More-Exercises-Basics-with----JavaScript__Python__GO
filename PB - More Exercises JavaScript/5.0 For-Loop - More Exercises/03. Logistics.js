function logistic(arg) {
    arg = [...arg.map(Number)]
    let cargoCount = arg.shift()

    const cargosObj = {'all_cargo': 0, 'bus': 0, 'truck': 0, 'train': 0}

    for (let i = 0; i < cargoCount; i++) {
        let currentCargo = arg[i]
        if (currentCargo <= 3) {
            cargosObj['bus'] += currentCargo
        } else if (4 <= currentCargo && currentCargo <= 11) {
            cargosObj['truck'] += currentCargo
        } else {
            cargosObj['train'] += currentCargo
        }

        cargosObj['all_cargo'] += currentCargo
    }

    let averagePrice = (cargosObj['bus'] * 200 + cargosObj['truck'] * 175 + cargosObj['train'] * 120) / cargosObj['all_cargo']

    console.log(`${averagePrice.toFixed(2)}`)
    console.log(`${(cargosObj['bus'] /   cargosObj['all_cargo'] *   100).toFixed(2)}%`)
    console.log(`${(cargosObj['truck'] / cargosObj['all_cargo'] * 100).toFixed(2)}%`)
    console.log(`${(cargosObj['train'] / cargosObj['all_cargo'] * 100).toFixed(2)}%`)
}

// logistic([4, 1, 5, 16, 3, ])