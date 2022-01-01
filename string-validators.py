if __name__ == '__main__':
    s = input()
    isalnum = False
    isalpha = False
    isdigit = False
    islower = False
    isupper = False
    for i in range(len(s)):
        if s[i].isalnum():
            isalnum = True
        if s[i].isalpha():
            isalpha = True  
        if s[i].isdigit():
            isdigit = True
        if s[i].islower():
            islower = True
        if s[i].isupper():
            isupper = True              

    print(isalnum)
    print(isalpha)
    print(isdigit)
    print(islower)
    print(isupper)
