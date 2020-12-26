x1, y1, x2, y2, x, y = [float(input()) for _ in range(6)]

if (x == x1 and y1 <= y <= y2) or (x == x2 and y1 <= y <= y2) or \
    (y == y1 and x1 <= x <= x2) or (y == y2 and x1 <= x <= x2):
    print("Border")
elif y == y1 and y == y2 and x1 < x < x2:
    print("Border")
else:
    print("Inside / Outside")


'''
2
-3
12
3
8
-1
'''
