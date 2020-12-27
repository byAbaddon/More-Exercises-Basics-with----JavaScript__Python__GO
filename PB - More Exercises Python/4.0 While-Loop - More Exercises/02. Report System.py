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
