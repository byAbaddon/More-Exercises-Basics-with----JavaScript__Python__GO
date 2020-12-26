season = input()
km_per_month = float(input())

tariff_type = ''
seasons_price_table_dict = {
    'min_tariff' : {'Autumn' : 0.75, 'Spring' : 0.75, 'Summer' : 0.9, 'Winter' : 1.05},
    'middle_tariff' : {'Autumn' : 0.95,'Spring' : 0.95, 'Summer' : 1.10, 'Winter' : 1.25},
    'max_tariff' : {'Autumn' : 1.45,'Spring' : 1.45, 'Summer' : 1.45, 'Winter' : 1.45},
}

if km_per_month <= 5000:
    tariff_type = 'min_tariff'
elif 5000 < km_per_month <= 10000:
    tariff_type = 'middle_tariff'
elif 10000 < km_per_month <= 20000:
    tariff_type = 'max_tariff'

count_mount = 4
tax = 0.90
money_passed_km = (km_per_month * seasons_price_table_dict[tariff_type][season]) * count_mount * tax

print(f'{money_passed_km:.2f}')


'''
Winter
16042
'''
