num_junior_biker, num_senior_biker, route_type = int(input()) , int(input()), input()
group_map = {
    'juniors' : {'trail' : 5.50, 'cross-country' : 8.00, 'downhill' : 12.25, 'road' : 20.00},
    'seniors' : {'trail' : 7.00, 'cross-country' : 9.50, 'downhill' : 13.75, 'road' : 21.50},
}

juniors_money = group_map['juniors'][route_type] * num_junior_biker
seniors_money = group_map['seniors'][route_type] * num_senior_biker
all_bikers_money = juniors_money + seniors_money
biker_count = num_junior_biker + num_senior_biker

if route_type == 'cross-country' and biker_count >= 50:
    all_bikers_money = all_bikers_money * 0.75

all_bikers_money = all_bikers_money * 0.95

print(f'{all_bikers_money + 0.001:.2f}')


'''
10
20
trail
'''
