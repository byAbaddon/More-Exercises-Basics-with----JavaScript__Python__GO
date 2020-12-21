function pointOnRectangleBorder(...arg) {
    let [x1, y1, x2, y2, x, y] = [...arg.map(Number)]

    if ((x == x1 && y1 <= y && y <= y2) || (x == x2 && y1 <= y && y <= y2) ||
        (y == y1 && x1 <= x && x <= x2) || (y == y2 && x1 <= x && x <= x2)) {
        return 'Border'
    } else if (y == y1 && y == y2 && x1 < x && x < x2) {
        return 'Border'
    } else {
        return 'Inside / Outside'
    }
}

//console.log(pointOnRectangleBorder(2, -3, 12, 3, 1, 3)) //Inside / Outside
