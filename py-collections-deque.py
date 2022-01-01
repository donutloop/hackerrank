

from collections import deque

d = deque()

commandCount = int(input())

for i in range(commandCount):
    values = input().split()
    method = getattr(d, values[0])
    if len(values) == 2:
        method(int(values[1]))
    else:
        method()        

t = ""
for n in d:
    t += str(n) + " "

print(t)    
