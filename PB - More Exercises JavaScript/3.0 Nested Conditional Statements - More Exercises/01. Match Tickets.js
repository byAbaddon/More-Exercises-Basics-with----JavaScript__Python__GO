function matchTickets(...arg) {
    let [budget, ticket, people] = [...arg]
    const tickets_type = {'VIP': 499.99,'Normal': 249.99 }
        
    if (1 <= people && people <= 4){
        budget *= 0.25
    } else if (5 <= people && people <= 9) {
        budget *= 0.40
    } else if (10 <= people && people <= 24) {
        budget *= 0.50
    } else if (25 <= people && people <= 49) {
        budget *= 0.60
    } else {
        budget *= 0.75
    }

    let ticket_price = tickets_type[ticket] * people
    let money_left = budget - ticket_price

    if (money_left > 0)
        return `Yes! You have ${money_left.toFixed(2)} leva left.`
    return `Not enough money! You need ${Math.abs(money_left).toFixed(2)} leva.`
}

// console.log(matchTickets(1000, 'Normal', 1))     //Yes! You have 0.01 leva left.