from itertools import product

start, end, ignore = [ord(input()) for _ in range(3)]
lst = map(chr, range(start , end + 1))
result = [x for x in map(''.join, product(lst, repeat=3)) if chr(ignore) not in x]
print(*result, len(result))

#-------------------------------------------------------------(2)---------------

l_one, l_two, l_ignore = [input() for _ in range(3)]
collection_list = []

for i in range(ord(l_one),ord(l_two) + 1):
    for j in range(ord(l_one),ord(l_two) + 1):
        for k in range(ord(l_one),ord(l_two) + 1):
            current = chr(i) + chr(j) + chr(k)
            if l_ignore not in current:
                collection_list.append(current)

print(' '.join(collection_list), len(collection_list))

'''
a
c
b
'''
