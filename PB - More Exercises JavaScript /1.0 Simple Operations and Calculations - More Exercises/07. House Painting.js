function housePainting(...arg) {
    let [x, y, h] = [...arg.map(Number)]
    let doorArea = 2 * 1.2
    let windowsArea = 2 * (1.5 * 1.5)
    let frontWallsArea = 2 * (x * x) - doorArea
    let sideWalls = 2 * (x * y) - windowsArea
    let totalWallArea = frontWallsArea + sideWalls
    let greenPaint = totalWallArea / 3.4
    let roof_square_area = 2 * (x * y)
    let roofTriangleArea = (x * h)
    let totalRoofArea = roof_square_area + roofTriangleArea
    let redPaint = totalRoofArea / 4.3

    return `${greenPaint.toFixed(2)}\n${redPaint.toFixed(2)}`
}

//console.log(housePainting(6, 10, 5.2)) //54.44 / 35.16