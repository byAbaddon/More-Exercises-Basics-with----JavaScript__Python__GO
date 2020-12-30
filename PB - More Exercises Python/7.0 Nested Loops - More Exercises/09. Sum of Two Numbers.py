start, end, magic_num = [int(input()) for _ in range(3)]
counter = 0

for a in range(start, end + 1 ):
    for b in range(start, end + 1):
        counter += 1
        if a + b == magic_num:
            print(f'Combination N:{counter} ({a} + {b} = {a + b})')
            exit()

print(f'{counter} combinations - neither equals {magic_num}')


'''
23
24
20
'''
