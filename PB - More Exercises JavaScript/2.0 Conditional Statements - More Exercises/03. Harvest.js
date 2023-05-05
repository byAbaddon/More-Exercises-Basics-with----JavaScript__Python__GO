function harvest([vineyard, kvM, necessaryWine, workers]) {
    let allWine = vineyard * 0.40 * kvM / 2.5
    let bonusWine = allWine - necessaryWine
    let wineForWorkers = bonusWine / workers
 
    if (necessaryWine > allWine)
        return `It will be a tough winter! More ${Math.abs(~~bonusWine)} liters wine needed.`
    return `Good harvest this year! Total wine: ${~~allWine} liters.\n${Math.ceil(bonusWine)} liters left -> ${Math.ceil(wineForWorkers)} liters per person.`
}

//-----------------------------------------------(2)  //old input no't array


function harvest(...arg) {
    let [vineyard, kvM, necessaryWine, workers] = [...arg.map(Number)]
    let allWine = (vineyard *= 0.40) * kvM / 2.5
    let bonusWine = allWine - necessaryWine
    let wineForWorkers = bonusWine / workers

    if (necessaryWine > allWine)
        return `It will be a tough winter! More ${Math.abs(Math.trunc(bonusWine))} liters wine needed.`
    return `Good harvest this year! Total wine: ${~~allWine} liters.\n${Math.ceil(bonusWine)} liters left -> ${Math.ceil(wineForWorkers)} liters per person.`
}

// console.log(harvest(650, 2, 175, 3))
