n = int(input())

[print(f'{"*" * (2 * n)}{" " * n }{"*" * (2 * n)}')for _ in range(1)]

for row in range(n - 2):
    if row == (n - 1) // 2 - 1: element = "|" * n
    else: element = " " * n
    print(f'*{"/" * (n * 2 - 2)}*{element}*{"/" * (n * 2 - 2)}*')

[print(f'{"*" * (2 * n)}{" " * n }{"*" * (2 * n)}')for _ in range(1)]

