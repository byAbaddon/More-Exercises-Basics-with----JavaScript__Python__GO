function weatherForecastTwo(temp) {
    if (5 <= temp && temp <= 11.9) return 'Cold'
    if (12 <= temp && temp <= 14.9) return 'Cool'
    if (15 <= temp && temp <= 20) return 'Mild'
    if (20.1 <= temp && temp <= 25.9) return 'Warm'
    if (26 <= temp && temp <= 35) return 'Hot'
    return 'unknown'
}

// console.log(weatherForecastTwo(8)) 'Cold'