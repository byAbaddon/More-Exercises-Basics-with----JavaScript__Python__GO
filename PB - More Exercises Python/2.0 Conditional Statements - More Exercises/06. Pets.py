from math import ceil, floor

days, kg_food, dog_food_kg, cat_food_kg, turtle_food_kg = [float(input())  for _ in range(5)]
food = (days * dog_food_kg) + (days * cat_food_kg) + (days * turtle_food_kg) / 1000

if food <= kg_food:
    print(f'{floor(kg_food - food)} kilos of food left.')
else:
    print(f'{ceil(food - kg_food)} more kilos of food are needed.')


'''
5
10
2.1
0.8
321
'''
