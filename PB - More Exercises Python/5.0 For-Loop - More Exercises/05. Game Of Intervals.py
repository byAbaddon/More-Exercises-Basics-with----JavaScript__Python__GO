from collections import OrderedDict
move = int(input())
points = 0

percent_dict = OrderedDict()
percent_dict['From 0 to 9:'] = 0
percent_dict['From 10 to 19:'] = 0
percent_dict['From 20 to 29:'] = 0
percent_dict['From 30 to 39:'] = 0
percent_dict['From 40 to 50:'] = 0
percent_dict['Invalid numbers:'] = 0


for _ in range(move):
    num = int(input())
    if  0 <= num <= 9:
        points += num * 0.20
        percent_dict['From 0 to 9:'] += 1
    elif 10 <= num <= 19:
        points += num * 0.30
        percent_dict['From 10 to 19:'] += 1
    elif 20 <= num <= 29:
        points += num * 0.40
        percent_dict['From 20 to 29:'] += 1
    elif 30 <= num <= 39:
        points += 50
        percent_dict['From 30 to 39:'] += 1
    elif 40 <= num <= 50:
        points += 100
        percent_dict['From 40 to 50:'] += 1
    else:
        points /= 2
        percent_dict['Invalid numbers:'] += 1
        
print(f'{points:.2f}')

for k,v in percent_dict.items():
    print(f'{k} {v / move * 100:.2f}%')





'''
10
43
57
-12
23
12
0
50
40
30
20
----------------------
3
12
-23
46
'''



