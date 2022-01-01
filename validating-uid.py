# It must contain at least  2 uppercase English alphabet characters.
# It must contain at least 3  digits (0 - 9).
#It should only contain alphanumeric characters (a - z ,  A - Z &  0-9 ).
# No character should repeat.
#There must be exactly 10 characters in a valid UID.

def validateUID(s):
    
    if len(s) != 10:
        return False
    
    upperCaseCount = 0
    digitCount = 0
    chars = dict() 
    for char in s:
        if char in chars:
            return False
        
        if ord(char) >= 65 and ord(char) <= 90:    
            upperCaseCount = upperCaseCount + 1   
        elif ord(char) >= 97 and ord(char) <= 122:
            pass
        elif ord(char) >= 48 and ord(char) <= 57:   
            digitCount = digitCount + 1
        else: 
        
            return False
        chars[char] = True
        
    return digitCount >= 3 and upperCaseCount >= 2

if __name__ == '__main__':
    n = int(input())
    
    for i in range(n):
        s = input()
        ok = validateUID(s)
        if ok:
            print("Valid")
        else:
            print("Invalid")    
        
    
    
