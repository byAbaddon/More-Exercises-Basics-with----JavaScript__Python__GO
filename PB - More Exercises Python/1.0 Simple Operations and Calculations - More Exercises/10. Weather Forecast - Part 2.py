def weather(temp):
    return 'Cold' if 5 <= temp <= 11.9 else \
           'Cool' if 12 <= temp <= 14.9 else \
           'Mild' if 15 <= temp <= 20 else \
           'Warm' if 20.1 <= temp <= 25.9 else \
           'Hot' if 26 <= temp <= 35 else \
           'unknown'

print(weather(float(input())))


#8
