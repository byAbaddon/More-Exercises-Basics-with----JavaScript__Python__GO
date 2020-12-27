chars = input()
word_collection = ''
word = ''
c, o, n = False, False, False

while chars != 'End':
    if chars.isalpha():
       if (chars =='c' and c != True) or (chars =='o' and o != True) or (chars =='n' and n != True):
            if chars =='c': c = True
            if chars =='o': o = True
            if chars =='n': n = True
       else:
          word += chars

    if c and o and n:
       word_collection += word + ' '
       c, o, n = False, False, False
       word = ''

    chars =  input()

print(word_collection)


'''
H
n
e
l
l
o
o
c
t
c
h
o
e
r
e
n
e
End
'''
