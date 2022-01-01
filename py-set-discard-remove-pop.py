n = int(input())
s = set(map(int, input().split()))
commandCount = int(input())

for i in range(commandCount):
    values = input().split()
    method = getattr(s, values[0])
    if len(values) == 2:
        method(int(values[1]))
    else:
        method()        

ns = 0
for n in s:
    ns += n
        
print(ns)    
