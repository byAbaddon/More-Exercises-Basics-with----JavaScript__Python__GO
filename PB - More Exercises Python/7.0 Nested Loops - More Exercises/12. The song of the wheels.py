ctrl_num = int(input())
pass_count = 0
password = ''

for a in range(1, 10):
    for b in range(1, 10):
        for c in range(1, 10):
            for d in range(1, 10):
                if ctrl_num == a * b + c * d  and (a < b and c > d):
                    print(f'{a}{b}{c}{d}', end=' ')
                    pass_count += 1
                    if pass_count == 4:
                        password = f'{a}{b}{c}{d}'

if pass_count > 0:
   print()
if password:
    print('Password:', password)
else:
    print('No!')


'''
11
----- or
139
'''
