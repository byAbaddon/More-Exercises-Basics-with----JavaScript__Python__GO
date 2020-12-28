months_count = int(input())
all_electricity = sum([float(input()) for _ in range(months_count)])
water = 20 * months_count
internet = 15 * months_count
other = all_electricity * 1.20  + water * 1.20  + internet * 1.2
average = (all_electricity + water + internet + other) / months_count

print(f'Electricity: {all_electricity:.2f} lv')
print(f'Water: {water:.2f} lv')
print(f'Internet: {internet:.2f} lv')
print(f'Other: {other:.2f} lv')
print(f'Average: {average:.2f} lv')


'''
5
68.63
89.25
132.53
93.53
63.22
'''
