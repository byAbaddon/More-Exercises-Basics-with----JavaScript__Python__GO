stadium = int(input())
fens = int(input())
sectors_list = [input() for _ in range(fens)]

print(f'{sectors_list.count("A") / fens * 100:.2f}%')
print(f'{sectors_list.count("B") / fens * 100:.2f}%')
print(f'{sectors_list.count("V") / fens * 100:.2f}%')
print(f'{sectors_list.count("G") / fens * 100:.2f}%')
print(f'{fens / stadium * 100:.2f}%')

'''
76
10
A
V
V
V
G
B
A
V
B
B
'''
