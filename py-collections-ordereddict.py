from collections import OrderedDict

if __name__ == '__main__':
    a = int(input().strip())
    basket = OrderedDict()
    for n in range(a):
        values = input().strip().split()
        productNetPrice = int(values[len(values)-1])
        values = values[:len(values)-1]
        productName = " ".join(values)
        if productName in basket:
           basket[productName] += productNetPrice
        else:        
           basket[productName] = productNetPrice 
           
           
    for key, value in basket.items():
        print(key, value)
