budget, season = float(input()), input()
class_type, car_type = '', ''

if budget <= 100:
    class_type = 'Economy class'
    if season == 'Summer':
        budget *= 0.35
        car_type = 'Cabrio'
    else:
        budget *= 0.65
        car_type = 'Jeep'
elif 100 < budget <= 500:
    class_type = 'Compact class'
    if season == 'Summer':
        budget *= 0.45
        car_type = 'Cabrio'
    else:
        budget *= 0.80
        car_type = 'Jeep'
else:
    class_type = 'Luxury class'
    budget *= 0.90
    car_type = 'Jeep'

print(f'{class_type}\n{car_type} - {budget:.2f}')


'''
70.50
Winter
'''
