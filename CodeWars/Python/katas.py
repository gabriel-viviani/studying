# make a function that can take any non-negative integer as a argument and return it with its digits in descending order

def Descending_Order(num):
    return int("".join(sorted(str(num), reverse=True)))