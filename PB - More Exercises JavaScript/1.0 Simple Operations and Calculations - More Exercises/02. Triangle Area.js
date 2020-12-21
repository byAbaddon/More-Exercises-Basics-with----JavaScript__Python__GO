function triangleArea(...arg) {
    let [a, h] = [...arg.map(Number)]
    return (a * h / 2).toFixed(2)
}

//console.log(triangleArea(20, 30))  //300.00


