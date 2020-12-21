function bikeRace(numJuniorBiker, numSeniorBiker, routeType) {
    numJuniorBiker = Number(numJuniorBiker)
    numSeniorBiker = Number(numSeniorBiker)
    const group_map = {
        'juniors' : {'trail' : 5.50, 'cross-country' : 8.00, 'downhill' : 12.25, 'road' : 20.00},
        'seniors' : {'trail' : 7.00, 'cross-country' : 9.50, 'downhill' : 13.75, 'road' : 21.50},
    }
    
    let juniors_money = group_map['juniors'][routeType] * numJuniorBiker
    let seniors_money = group_map['seniors'][routeType] * numSeniorBiker
    let all_bikers_money = juniors_money + seniors_money
    let biker_count = numJuniorBiker + numSeniorBiker
    
    if (routeType == 'cross-country' && biker_count >= 50){
        all_bikers_money = all_bikers_money * 0.75
    }

    all_bikers_money = all_bikers_money * 0.95
    
    return `${(all_bikers_money + 0.001).toFixed(2)}`
}

// console.log(bikeRace(10, 20, 'trail'))   //185.25
//  console.log(bikeRace(20, 25, 'cross-country'))   //377.63





