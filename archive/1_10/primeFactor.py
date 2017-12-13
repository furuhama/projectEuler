import math

def get_prime_factor(num):
    num = int(num)
    if num == 1:
        print('end')
    else:
        for i in range(num):
            if i == 0:
                next
            elif num % (i + 1) == 0:
                print(i + 1)
                break
            else:
                next
        get_prime_factor(num / (i + 1))

if __name__ == '__main__' :
    n = int(input())

    get_prime_factor(n)

