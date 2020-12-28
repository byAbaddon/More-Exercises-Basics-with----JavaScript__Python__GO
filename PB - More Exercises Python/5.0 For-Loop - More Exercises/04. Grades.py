students_count = int(input())
grades = [float(input()) for _ in range(students_count)]

def calc_grades(*args):
    test = 0
    if args[0] == 5:
        test = lambda x: x >= args[0]
    elif args[0] == 4:
        test = lambda x: args[0] <= x <= args[1]
    elif args[0] == 3:
        test = lambda x: args[0] <= x <= args[1]
    elif args[0] < 3:
        test = lambda x: args[0] > x

    return f'{len(list(filter(test, grades))) / students_count * 100:.2f}%'


print(f'Top students: {calc_grades(5)}')
print(f'Between 4.00 and 4.99: {calc_grades(4, 4.99)}')
print(f'Between 3.00 and 3.99: {calc_grades(3, 3.99)}')
print(f'Fail: {calc_grades(2.99)}')
print(f'Average: {sum(grades) / students_count:.2f}')


'''
6
2
3
4
5
6
2.2
'''
