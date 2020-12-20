function trainingLab(...arg) {
    let [height, width] = [...arg.map(Number)]
    let hallway = 100
    let lostWorkPosition = 3
    let workPositionWidth = 70
    let workPositionHeight = 120

    height *= 100
    width *= 100
    width -= hallway

    let rest = width % workPositionWidth
    width -= rest
    let bureauCount = width / workPositionWidth

    rest = height % workPositionHeight
    height -= rest
    let rowCount = height / workPositionHeight

    return bureauCount * rowCount - lostWorkPosition
}

// console.log(trainingLab(15, 8.9))  //129