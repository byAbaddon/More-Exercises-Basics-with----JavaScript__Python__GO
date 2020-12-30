start, end = int(input()), int(input()) + 1

for a in range(start, end):
    for b in range(start, end):
        for c in range(start, end):
            for d in range(start, end):
                if not a & 1 and d & 1 or a & 1 and not d & 1:
                    if a > d:
                        if (b + c) % 2 == 0:
                            print(str(a) + str(b) + str(c) + str(d), end=' ')
