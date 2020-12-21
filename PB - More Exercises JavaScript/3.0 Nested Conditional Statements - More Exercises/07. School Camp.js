function schoolCamp(...arg) {
    let [season, gender, count_students, count_nights] = [...arg.filter(el => isNaN(el) ? el : Number(el))]

    const holiday_type_dict = {
        'girls' : {'Winter' : 'Gymnastics', 'Spring' : 'Athletics', 'Summer' : 'Volleyball'},
        'boys' : {'Winter' : 'Judo', 'Spring' : 'Tennis', 'Summer' : 'Football'},
        'mixed' : {'Winter' : 'Ski', 'Spring' : 'Cycling', 'Summer' : 'Swimming'},
    }
    
    const price_type_dict = {
        'girls' : {'Winter' : 9.60, 'Spring' : 7.20, 'Summer' : 15.00},
        'boys' : {'Winter' : 9.60, 'Spring' : 7.20, 'Summer' : 15.00},
        'mixed' : {'Winter' : 10.00, 'Spring' : 9.50, 'Summer' : 20.00},
    }

    let sport_type = holiday_type_dict[gender][season]
    let total_price = price_type_dict[gender][season] * count_students * count_nights

    if (10 <= count_students && count_students < 20) {
        total_price *= 0.95
    } else if (20 <= count_students && count_students < 50) {
        total_price *= 0.85
    } else if (count_students >= 50) {
        total_price *= 0.50
    }

    return `${sport_type} ${total_price.toFixed(2)} lv.`
}

//console.log(schoolCamp('Spring', 'girls', 20, 7))   //Athletics 856.80 lv.
