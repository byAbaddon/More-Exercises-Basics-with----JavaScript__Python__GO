function fishLand(...arg) {
    let [skumriqPrice, cacaPrice, kgPalamud, kgSafrid, kgMidi] = [...arg.map(Number)]
    let palamudPrice = skumriqPrice + 0.6 * skumriqPrice
    let safridPrice = cacaPrice + 0.8 * cacaPrice
    let price = (kgPalamud * palamudPrice + kgSafrid * safridPrice + kgMidi * 7.5)
    return price.toFixed(2)
}

// console.log(fishLand(6.90, 4.20, 1.5, 2.5, 1))  //42.96