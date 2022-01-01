from itertools import product

a = map(int, list(input().split()))
b = map(int,list(input().split()))

for c in list(product(a, b)):
    print(c, end=" ")
