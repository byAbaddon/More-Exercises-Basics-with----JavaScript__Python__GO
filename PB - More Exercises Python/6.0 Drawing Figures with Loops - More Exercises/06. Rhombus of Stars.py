n = int(input())
[print(f'{" " * (n - r) }{"* " * r}')for r in range(1, n + 1)]
[print(f'{" " * r }{"* " * (n - r)}')for r in range(1, n)]

