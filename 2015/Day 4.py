# Part 1

print('Part 1')

import hashlib

input = "yzbqklnj"

testnum = 0
teststring = (input + str(testnum))


while not (hashlib.md5(teststring.encode('utf-8')).hexdigest()[:5] == "00000"):
    testnum += 1
    teststring = (input + str(testnum))

print('hash = ' + hashlib.md5(teststring.encode('utf-8')).hexdigest())
print('number = ' + str(testnum))

# Part 2

print('Part 2')

import hashlib

input = "yzbqklnj"

testnum = 0
teststring = (input + str(testnum))


while not (hashlib.md5(teststring.encode('utf-8')).hexdigest()[:6] == "000000"):
    testnum += 1
    teststring = (input + str(testnum))

print('hash = ' + hashlib.md5(teststring.encode('utf-8')).hexdigest())
print('number = ' + str(testnum))