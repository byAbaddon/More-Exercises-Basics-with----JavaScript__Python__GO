height = float(input()) * 100
width = float(input()) * 100 - 100
rest = width % 70
width -= rest
bureauCount = width / 70
rest = height % 120
height -= rest
rowCount = height / 120
print(bureauCount * rowCount - 3)

#---------------------------------------------(2)---------------------------------------
print(int(float(input()) * 100 / 120) * int((float(input()) * 100 - 100) / 70) - 3)

'''
15
8.4
---------
15
8.9
'''   #129



