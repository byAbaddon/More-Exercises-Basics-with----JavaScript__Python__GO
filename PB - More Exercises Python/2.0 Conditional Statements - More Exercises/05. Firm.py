from math import trunc

hours, day, workers = [int(input()) for _ in range(3)]

remnant_time = trunc(day * 0.90 * 8)
extra_work_time = workers * (2 * day)
total_time = remnant_time + extra_work_time

if total_time >= hours:
    print(f'Yes!{total_time - hours} hours left.')
else:
    print(f'Not enough time!{hours - total_time} hours needed.')


'''
90
7
3
'''
