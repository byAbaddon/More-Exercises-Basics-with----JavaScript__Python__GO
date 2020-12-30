days, hours = [int(input()) + 1 for _ in range(2)]
current_tax, total_sum = 0, 0

for day in range(1, days):
    for hour in range(1, hours):
        if not day & 1 and hour & 1:
            current_tax += 2.50
        elif day & 1 and not hour & 1:
            current_tax += 1.25
        else:
            current_tax += 1.00
    print(f'Day: {day} - {current_tax:.2f} leva')
    total_sum += current_tax
    current_tax = 0

print(f'Total: {total_sum:.2f} leva')


'''
2
5
'''
