if __name__ == '__main__':
    a = int(input().strip())
    alist = input().strip().split()
    b = int(input().strip())
    blist = input().strip().split()
    
    print(len(set(alist).union(blist)))
