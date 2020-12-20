type_fuel, quantity, club_card = input(), float(input()), input()
fuel_price_dict = {'Gasoline' : 2.22,'Diesel' : 2.33,'Gas' : 0.93}
discount = {'Gasoline' : 0.18,'Diesel' : 0.12,'Gas' : 0.08}
subtotal_price, total_price = 0

if club_card == 'Yes':
    subtotal_price = fuel_price_dict[type_fuel] - discount[type_fuel]
    total_price = subtotal_price * quantity
else:
    total_price = fuel_price_dict[type_fuel] * quantity

if 20 <= quantity <= 25:
    total_price *= 0.92
elif quantity > 25:
     total_price *= 0.90

print(f'{total_price:.2f} lv.')



'''
Gas
30
Yes
'''
