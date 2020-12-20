function circleAreaAndPerimeter(n) {
    let face = Math.PI * n * n
    let prm =  Math.PI * 2 * n
    return `${face.toFixed(2)}\n${prm.toFixed(2)}`  
}

// console.log(circleAreaAndPerimeter(3)) //28.27   / 18.85