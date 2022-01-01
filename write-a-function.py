def is_leap(year):
    if year % 4 == 0:
        return year % 100 != 0 or year % 400 == 0    
    else:
        return False

