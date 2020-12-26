budget, season = float(input()), input()
location, place_type = '', ''

if budget <= 1000:
    place_type = 'Camp'
    if season == 'Summer':
        location = 'Alaska'
        budget *= 0.65
    else:
        location = 'Morocco'
        budget *= 0.45
elif 1000 < budget <= 3000:
    place_type = 'Hut'
    if season == 'Summer':
        location = 'Alaska'
        budget *= 0.80
    else:
        location = 'Morocco'
        budget *= 0.60
else:
    place_type = 'Hotel'
    if season == 'Summer':
        location = 'Alaska'
        budget *= 0.90
    else:
        location = 'Morocco'
        budget *= 0.90

print(f'{location} - {place_type} - {budget:.2f}')


'''
800
Summer
'''

