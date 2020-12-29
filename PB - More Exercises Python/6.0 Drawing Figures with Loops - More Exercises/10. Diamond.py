n = int(input())

if n == 1:
    print('*')
elif n == 2:
    print('**')
else:
    if n & 1:
        element = "*"
    else:
        element = "**"
#top row
    [print(f'{"-" * ((n - 1) // 2 )}{element}{"-" * ((n - 1) // 2 )}') for row in range(1)]

#---------------------------------------------------
#middle top
    if n > 4 and n & 1: #up  /odd
       [print(f"{'-' * (n // 2 - i)}*{'-' * ((i + i) -1)}*{'-' * (n // 2 - i)}") for i in range(1, n // 2)]
    else: #even
        [print(f"{'-' * ((n // 2 - i) - 1)}*{'-' * (i + i)}*{'-' * ((n // 2 - i)-1)}") for i in range(1, n // 2 - 1)]

# middle line
    [print(f"*{'-' * (n -2)}*") for _ in range(1)]

#under middle
    if n > 4 and n & 1:
        [print(f"{'-' * (n // 2 - i)}*{'-' * (i + i - 1)}*{'-' * (n // 2 - i)}") for i in range(n // 2 -1, 0, -1)]
    else:
         #even
        [print(f"{'-' * (n // 2 - i)}*{'-' * (i + i -2)}*{'-' * (n // 2 - i)}") for i in range(n // 2 -1 , 1, -1)]
#---------------------------------------------------
#buttom row
    [print(f'{"-" * ((n - 1) // 2 )}{element}{"-" * ((n - 1) // 2 )}') for _ in range(1)]
    
    
    
#----------------------------------------------------------------------(2)------------------------------------------------------------------
'''
n = int(input())

leftRight = int((n - 1) / 2)
# 1st row
for i in range(1, leftRight + 1):
    print(f"{'-' * leftRight}*", end="")
    mid = n - 2 * leftRight - 2
    if mid >= 0:
        print(f"{'-' * mid}*", end="")

    print(f"{'-' * leftRight}")
    leftRight -= 1

# Bottom part
for i in range(1, int(((n+1)/2) + 1)):
    leftRight = i - 1
    mid = n - 2 * leftRight - 2
    print(f"{'-' * leftRight}*", end="")
    if mid >= 0:
        print(f"{'-' * mid}*", end="")
    print(f"{'-' * leftRight}")
'''
    
