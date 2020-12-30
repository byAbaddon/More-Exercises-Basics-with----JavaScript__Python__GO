a, b, max_pass = [int(input()) for _ in range(3)]
code_a, code_b, password = 35, 64, ''

while code_a <= 55 and code_b <= 96:
    for x in range(1, a + 1):
        for y in range(1, b + 1):
            password += f'{chr(code_a)}{chr(code_b)}{x}{y}{chr(code_b)}{chr(code_a)}|'
            max_pass -= 1
            code_a += 1
            code_b += 1
            if code_a > 55: code_a = 35
            if code_b > 96: code_b = 64

            if max_pass == 0 or (x == a and y == b):
                print(password)
                exit()


'''
2
3
10
'''
