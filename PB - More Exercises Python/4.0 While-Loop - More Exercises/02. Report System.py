
target_sum = int(input())
statistics_obj = {'CS': {'money': 0, 'count': 0}, 'CC': {'money': 0, 'count': 0}}
current_sum = 0

for i in range(100):
    try:
        money = int(input())
        if i & 1:
            type_pay = 'CC'
        else:
            type_pay = 'CS'
        if money > 100 and type_pay == 'CS' or money < 10 and type_pay == 'CC':
            print('Error in transaction!')
        else:
            current_sum += money
            statistics_obj[type_pay]['money'] += money
            statistics_obj[type_pay]['count'] += 1
            print('Product sold!')
            if statistics_obj[type_pay]['money'] >= target_sum:
                break
    except:
        break

if current_sum < target_sum:
    print('Failed to collect required money for charity.')
else:
    for k, v in statistics_obj.items():
        print(f'Average {k}: {v["money"] / v["count"]:.2f}')

# --------------------------------------------------------------(2)-----------------------------------------------

necessary_money = float(input())
cash_payment_count = 0
cash_payment = 0
card_payment = 0
card_payment_count = 0
total_money = 0
is_payment_cash = True

while True:
    try:
        products_price = float(input())
    except:
        break

    if is_payment_cash:
        is_payment_cash = False
        if products_price > 100:
            print('Error in transaction!')
        else:
            cash_payment += products_price
            cash_payment_count += 1
            print('Product sold!')
            total_money += products_price
    else:
        is_payment_cash = True
        if products_price < 10:
            print('Error in transaction!')
        else:
            card_payment += products_price
            card_payment_count += 1
            print('Product sold!')
            total_money += products_price

    if total_money >= necessary_money:
       break

if total_money < necessary_money:
    print('Failed to collect required money for charity.')
else:
    average_cash_payment = cash_payment / cash_payment_count
    average_card_payment = card_payment / card_payment_count
    print(f"Average CS: {average_cash_payment:.2f}")
    print(f"Average CC: {average_card_payment:.2f}")




'''
500
120
8
63
256
78
317
600
86
150
98
227
End
'''
