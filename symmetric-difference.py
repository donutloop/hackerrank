if __name__ == '__main__':
    a = int(input().strip())
    alist = input().strip().split()
    alist = list(map(int, alist))
    b = int(input().strip())
    blist = input().strip().split()
    blist = list(map(int, blist))
    
    filteredA = set(alist).difference(blist)
    filteredB = set(blist).difference(alist)
    for n in sorted(list(filteredB) + list(filteredA)):
        print(n)
