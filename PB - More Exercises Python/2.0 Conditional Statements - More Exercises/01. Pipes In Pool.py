volume, pipe_one, pipe_two, hours = [float(input()) for _ in range(4)]
pipe_one *= hours
pipe_two *= hours
pool = pipe_one + pipe_two
pool_percent = pool / volume * 100
pipe_one_percent = pipe_one / pool * 100
pipe_two_percent = pipe_two / pool * 100

if pool <= volume:
    print(f'The pool is {pool_percent:.2f}% full. Pipe 1: {pipe_one_percent:.2f}%. Pipe 2: {pipe_two_percent:.2f}%.')
else:
    print(f'For ${hours:.2f} hours the pool overflows with {pool - volume:.2f} liters.')



'''
1000
100
120
3
'''
