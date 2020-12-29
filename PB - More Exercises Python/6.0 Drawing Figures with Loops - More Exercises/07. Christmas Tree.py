n = int(input())
[print(f'{" " * n} | {" " * n}')for r in range(1)]
[print(f'{" " * (n - r)}{"*" * r} | {"*" * r}')for r in range(1, n + 1)]
