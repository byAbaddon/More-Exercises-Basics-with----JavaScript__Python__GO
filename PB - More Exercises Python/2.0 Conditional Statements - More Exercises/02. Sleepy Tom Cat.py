from math import trunc

rest_days = int(input())
game_Tom_Min = 30000
owner_work = 63
owner_off = 127
year = (365 - rest_days)
year_to_min = (year * 63) + (rest_days * 127)
subtotal = abs(game_Tom_Min - year_to_min)
hours = trunc(subtotal / 60)
minute = subtotal % 60

if game_Tom_Min >= year_to_min:
    print(f'Tom sleeps well\n{hours} hours and {minute} minutes less for play')
else:
    print(f'Tom will run away\n{hours} hours and {minute} minutes more for play')

#20
