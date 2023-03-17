c, r, t, season, holiday = [input() if i > 2 else int(input()) for i in range(5)]
data_dict = {'s_s': {'c': c * 2, 'r': 4.1 * r, 't': 2.5 * t}, 'a_w': {'c': c * 3.75, 'r': 4.5 * r, 't': 4.15 * t} }

if season == 'Spring' or season == 'Summer':
    price = sum(data_dict['s_s'].values())
else:
    price = sum(data_dict['a_w'].values())

if holiday == 'Y': price *= 1.15
if season == 'Spring' and t > 7: price *= 0.95
if season == 'Winter' and r >= 10: price *= 0.90
if c + r + t > 20: price *= 0.8

print(f'{(price + 2):.2f}')


#--------------------------------------------------------------------(2)--------------------------------------


chrysanthemums , roses, tulips = [int(input()) for _ in range(3)]
season, is_celebrate_day = [input() for _ in range(2)]

chrysanthemum_total, rose_total, tulip_total = 0, 0, 0

if season == 'Spring' or season == 'Summer':
    chrysanthemum_total = chrysanthemums * 2
    rose_total = roses * 4.1
    tulip_total = tulips * 2.5
elif season == 'Autumn' or season == 'Winter':
    chrysanthemum_total = chrysanthemums * 3.75
    rose_total = roses * 4.5
    tulip_total = tulips * 4.15

price = chrysanthemum_total + rose_total + tulip_total

if is_celebrate_day == 'Y':
    price  = price * 1.15

if season == 'Spring' and tulips > 7:
    price *= 0.95

if season == 'Winter' and roses >= 10:
    price *= 0.90

if chrysanthemums + roses + tulips> 20:
    price *= 0.8

print(f'{(price + 2):.2f}')


'''
2
4
8
Spring
Y
'''
