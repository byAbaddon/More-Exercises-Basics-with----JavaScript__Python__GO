detergent = int(input()) * 750
dishes, pots, count = 0, 0, 0
while True:
    try:
        current = int(input())
        count += 1
    except:
        break
    else:
        if count % 3 != 0:
          detergent -= current * 5
          dishes += current
        else:
          detergent -= current * 15
          pots += current

if detergent >= 0:
    print(f'Detergent was enough!\n{dishes} dishes and {pots} pots were washed.\nLeftover detergent {detergent} ml.')
else:
    print(f'Not enough detergent, {abs(detergent)} ml. more necessary!')
