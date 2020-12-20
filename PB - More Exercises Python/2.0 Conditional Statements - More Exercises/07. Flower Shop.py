from math import ceil, floor

count_mgn, count_zum, count_roses, count_cactus, gift_price = [float(input()) for _ in range(5)]
tax = 0.95
total_sum = (count_mgn * 3.25 + count_zum * 4 + count_roses * 3.5 + count_cactus * 8) * tax

if total_sum < gift_price:
    print(f'She will have to borrow {ceil(gift_price - total_sum)} leva.')
else:
    print(f'She is left with {floor(total_sum - gift_price)} leva.')

'''
2
3
5
1
50
'''
