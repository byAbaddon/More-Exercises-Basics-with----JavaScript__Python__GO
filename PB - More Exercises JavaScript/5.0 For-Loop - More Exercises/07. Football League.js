function footballLeague(arg) {
    let stadium = Number(arg.shift())
    let fens = Number(arg.shift())
    let sectorsList = [...arg] 

    console.log(`${(sectorsList.filter(el => el == "A").length / fens * 100).toFixed(2)}%`)
    console.log(`${(sectorsList.filter(el => el == "B").length / fens * 100).toFixed(2)}%`)
    console.log(`${(sectorsList.filter(el => el == "V").length / fens * 100).toFixed(2)}%`)
    console.log(`${(sectorsList.filter(el => el == "G").length / fens * 100).toFixed(2)}%`)
    console.log(`${(fens / stadium * 100).toFixed(2)}%`)
}

// footballLeague([76, 10, 'A', 'V', 'V', 'V', 'G', 'B', 'A', 'V', 'B', 'B'])