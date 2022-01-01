if __name__ == '__main__':
    values = input().strip().split()
    x = values[0]
    k = int(values[1])
    e = input().strip().replace("x", x)
    y = int(eval(e))
    print(y == k)
