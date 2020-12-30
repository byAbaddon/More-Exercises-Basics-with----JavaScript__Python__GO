one_coins, two_coins, five_coins, all_sum = [int(input()) for _ in range(4)]

for one in range(one_coins + 1):
    for two in range(two_coins + 1):
        for five in range(five_coins + 1):
            if one * 1 + two * 2 + five * 5 == all_sum:
                print(f'{one} * 1 lv. + {two} * 2 lv. + {five} * 5 lv. = {all_sum} lv.')




'''
3
2
3
10
'''
