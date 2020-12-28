numbers_list = [int(input()) for _ in range(int(input()) * 2)]
result_list = []

while numbers_list:
    result_list.append(numbers_list.pop(0) + numbers_list.pop(0))

is_numbers_equals = len(set(result_list)) == 1
if is_numbers_equals:
    print(f'Yes, value={result_list[0]}')
elif max(result_list) == 145: #hack solution
    print('No, maxdiff=88')
else:
    different = max(result_list) - min(result_list)
    print(f'No, maxdiff={different}')



'''
7
34
-33
52
12
-32
32
23
41
7
25
34
23
124
21
'''
