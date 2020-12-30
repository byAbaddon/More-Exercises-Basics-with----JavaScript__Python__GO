sectors, row_count, seats_count = input(), int(input()), int(input())
even_seats, counter = 0, 0

for sector in range(ord('A'), ord(sectors) + 1):
    row_count += 1
    for row in range(1, row_count):
        if row % 2 == 0:
           even_seats = 2
        else:
            even_seats = 0
        for seats in range(1, seats_count + even_seats + 1):
            print(f'{chr(sector)}{row}{chr(96 + seats)}')
            counter += 1

print(counter)


'''
'B'
3
2
'''


