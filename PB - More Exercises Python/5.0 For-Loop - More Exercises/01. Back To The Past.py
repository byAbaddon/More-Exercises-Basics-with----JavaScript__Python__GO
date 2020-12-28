money = float(input())
years = int(input())
years_old = 17
for x in range(1800, years + 1):
    years_old +=1
    if not x & 1:
        money -= 12000
    else:
        money -= 12000 + years_old * 50

if money >= 0:
    print(f'Yes! He will live a carefree life and will have {money:.2f} dollars left.')
else:
    print(f'He will need {abs(money):.2f} dollars to survive.')



'''
50000
1802
'''
