km, day_type, price = int(input()), input(), 0

if km < 20:
    if day_type == 'day':
        price = 0.79 * km + 0.70
    else:
        price = 0.90 * km + 0.70
elif 20 <= km < 100:
    price = km * 0.09
else:
    price = km * 0.06

print(f'{price:.2f}')


'''
7
night
'''
