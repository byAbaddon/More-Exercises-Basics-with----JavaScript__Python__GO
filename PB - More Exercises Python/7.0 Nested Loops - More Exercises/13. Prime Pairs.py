f_start = int(input())
s_start = int(input())
f_end = int(input()) + f_start
s_end = int(input()) + s_start
prime_number_list = [2, 3, 5, 7, 11, 13, 17, 19 , 23 , 29 , 31 , 37 , 41 , 43 , 47 , 53 , 59 , 61 , 67 , 71 , 73 , 79 , 83 , 89 , 97 , 101]


for i in range(f_start, f_end + 1):
   for j in range(s_start, s_end + 1):
       if i in prime_number_list and j in prime_number_list:
          print(str(i) + str(j))


'''
10
20
5
5
'''
