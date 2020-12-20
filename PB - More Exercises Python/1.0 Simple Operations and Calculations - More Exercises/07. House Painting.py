x, y, h = [float(input()) for _ in range(3 )]
door_area = 2 * 1.2
windows_area = 2 * (1.5 * 1.5)
front_walls_area = 2 * (x * x) - door_area
side_walls = 2 * (x * y) - windows_area
total_wall_area = front_walls_area + side_walls
green_paint = total_wall_area / 3.4
roof_square_area = 2 * (x * y)
roof_triangle_area = (x * h)
total_roof_area = roof_square_area + roof_triangle_area
red_paint = total_roof_area / 4.3
print(f'{green_paint:.2f}')
print(f'{red_paint:.2f}')


'''
6
10
5.2
'''
