men, women, table = [int(input()) for _ in range(3)]
visited = []
for m in range(1, men + 1):
    for w in range(1, women + 1):
        if table != len(visited):
            visited.append(f'{m} <-> {w}')
        else:
            break

[print(f'({x})', end=' ') for x in visited]


'''
2
2
6
'''
