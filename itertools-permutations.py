from itertools import permutations

if __name__ == '__main__':
    arr = list(map(str, input().split()))
    arr = list(permutations(arr[0],int(arr[1])))
    res = []
    for elems in arr:
        s = ""
        for char in elems:
            s += char
        res.append(s)
    res.sort()
    
    for chars in res:
        print(chars)
