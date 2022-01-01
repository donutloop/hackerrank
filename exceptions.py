if __name__ == '__main__':
    n = int(input().strip())
    
    for x in range(n):
       values = input().strip().split()
       a = values[0]
       b = values[1]
       
       try:
        print(int(int(a)//int(b)))
       except ZeroDivisionError as e:
        print("Error Code:", e)
       except ValueError as e:
        print("Error Code:", e) 
