n = int(input())

if n == 2:
    print(f'**\n||')
else:
#draw top
    if n & 1: #draw up  even
        for row in range(1, n // 2 + 1):
            print(f"{'-' * (n // 2 + 1 - row)}{'*' * (row * 2 - 1)}{'-' * (n // 2 + 1 - row)}")
    else: #draw up odd
        for row in range(1, n // 2):
            print(f"{'-' * (n // 2 - row)}{'*' * (row * 2)}{'-' * (n // 2 - row)}")
#draw middle part
    print('*' * n )
#draw bottom
    for _ in range(n // 2):
        print(f"|{'*' * (n - 2)}|")

