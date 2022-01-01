def swap_case(s):
    s = list(s)
    for i in range(0, len(s)):
        if ord(s[i]) >= 97 and ord(s[i]) <= 122:
           s[i] = chr(ord(s[i]) - 32)
        elif ord(s[i]) >= 65 and ord(s[i]) <= 90:
           s[i] = chr(ord(s[i]) + 32)
    
    return "".join(s)

