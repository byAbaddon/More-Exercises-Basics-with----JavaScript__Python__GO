from math import ceil, floor, trunc

vineyard, kvM, necessary_wine, workers = [float(input()) for _ in range(4)]

all_wine = (vineyard - (vineyard * 0.60)) * kvM / 2.5
bonus_wine = all_wine - necessary_wine
wine_for_workers = bonus_wine / workers

if necessary_wine > all_wine:
    print(f'It will be a tough winter! More {abs(trunc(bonus_wine))} liters wine needed.')
else:
    print(f'Good harvest this year! Total wine: {floor(all_wine)} liters.\n\
{ceil(bonus_wine)} liters left -> {ceil(wine_for_workers)} liters per person.')


'''
650
2
175
3
----------
1020
1.5
425
4

'''
