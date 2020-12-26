season, gender = input(), input()
count_students, count_nights = int(input()), int(input())

holiday_type_dict = {
    'girls' : {'Winter' : 'Gymnastics', 'Spring' : 'Athletics', 'Summer' : 'Volleyball'},
    'boys' : {'Winter' : 'Judo', 'Spring' : 'Tennis', 'Summer' : 'Football'},
    'mixed' : {'Winter' : 'Ski', 'Spring' : 'Cycling', 'Summer' : 'Swimming'},
}

price_type_dict = {
    'girls' : {'Winter' : 9.60, 'Spring' : 7.20, 'Summer' : 15.00},
    'boys' : {'Winter' : 9.60, 'Spring' : 7.20, 'Summer' : 15.00},
    'mixed' : {'Winter' : 10.00, 'Spring' : 9.50, 'Summer' : 20.00},
}


sport_type = holiday_type_dict[gender][season]
total_price = price_type_dict[gender][season] * count_students * count_nights

#diccount
if 10 <= count_students < 20:
    total_price *= 0.95
elif 20 <= count_students < 50:
   total_price  *= 0.85
elif count_students >= 50:
    total_price  *= 0.50


print(f'{sport_type} {total_price:.2f} lv.')


'''
Spring
girls
20
7
'''
