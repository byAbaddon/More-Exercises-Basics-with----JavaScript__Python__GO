function weddingSeats(arg) {
    let [sectors, rowCount, seatsCount] = [...arg.map(el => isNaN(el) ? el : Number(el))]
    let [evenSeats, counter] = [0, 0]

    for (let sector = 65; sector <= String(sectors).charCodeAt(0); sector++) {
        rowCount++
        for (let row = 1; row < rowCount; row++) {
            if (!(row & 1)) {
                evenSeats = 2
            } else {
                evenSeats = 0
            }

            for (let seats = 1; seats < (seatsCount + evenSeats + 1); seats++) {
                console.log(`${String.fromCharCode(sector)}${row}${String.fromCharCode(96 + seats)}`)
                counter++
            }
        }
    }

    return counter
}

//  console.log(weddingSeats(['B', 3, 2]))