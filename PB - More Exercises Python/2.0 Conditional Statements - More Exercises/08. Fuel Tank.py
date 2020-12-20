type_fuel, quantity = input(), int(input())
lst = ['Diesel','Gasoline','Gas']

if type_fuel not in lst:
    print('Invalid fuel!')
elif quantity >= 25:
    print(f'You have enough {type_fuel.lower()}.')
elif quantity < 25:
   print(f'Fill your tank with {type_fuel.lower()}!')


'''
Diesel
10
'''
