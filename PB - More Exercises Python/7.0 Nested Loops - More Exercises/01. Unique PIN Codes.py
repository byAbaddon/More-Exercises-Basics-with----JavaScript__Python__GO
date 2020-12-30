one_max, two_max, three_max = [int(input()) + 1 for _ in range(3)]
prime_number_list = [2, 3, 5, 7]

for f_digit in range(2, one_max, 2):
    for s_digit in range(2, two_max):
        for t_digit in range(2, three_max, 2):
            if s_digit in prime_number_list:
                print(f_digit, s_digit, t_digit)


'''
3
5
5
'''
#--------------------------------------------------(2)----------------------------------
[print(f'{f} : {s} : {t}')
 for t in range(2, int(input()),2)

 for s in range(2, int(input()))
 if s == 3 or s == 5 or s == 7

 for f in range(2, int(input()), 2)]
