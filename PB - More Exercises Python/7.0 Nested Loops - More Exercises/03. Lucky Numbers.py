from itertools import product
divider = int(input())

for el in map(str,product(range(1,10), repeat = 4)):
    concat = ''.join(filter(lambda y: y.isdigit(), el))
    if sum(map(int, concat[:2])) == sum(map(int, concat[2:])):
        if divider %  sum(map(int, concat[:2])) == 0:
            print(concat, end=' ')



#---------------------------------------------------(2)-------------------------
# divider_num = int(input())
#
# for a in range(1, 10):
#     for b in range(1, 10):
#         for c in range(1, 10):
#             for d in range(1, 10):
#                 if a + b == c + d and divider_num % (c + d) == 0:
#                     print(str(a) + str(b) + str(c) + str(d), end=' ')
