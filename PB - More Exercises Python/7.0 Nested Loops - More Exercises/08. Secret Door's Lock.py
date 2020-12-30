hundred_num, double_num, single_num = [int(input()) + 1 for _ in range(3)]
prime_nums = [2, 3, 5, 7]

for h in range(2, hundred_num, 2):
    for d in range(2, double_num):
        for s in range(2, single_num, 2):
            if d in prime_nums:
                print(f'{h} {d} {s}')


'''
3
5
5
'''
