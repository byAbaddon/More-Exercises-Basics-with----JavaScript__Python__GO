budget, ticket, people = float(input()), input(), int(input())
tickets_type = {'VIP' : 499.99, 'Normal': 249.99}

if  1 <= people <= 4:
    budget *= 0.25
elif 5 <= people <= 9:
    budget *= 0.40
elif 10 <= people <= 24:
    budget *= 0.50
elif 25 <= people <= 49:
    budget *= 0.60
else:
    budget *= 0.75

ticket_price = tickets_type[ticket] * people
money_left = budget - ticket_price

if money_left > 0:
    print(f'Yes! You have {money_left:.2f} leva left.')
else:
    print(f'Not enough money! You need {abs(money_left):.2f} leva.')
