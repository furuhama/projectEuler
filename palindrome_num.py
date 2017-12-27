""" Project Euler Q.4 """

from math import sqrt


def main():
    if search_6digits_number() != 0:
        print(search_6digits_number())
    elif search_5digits_number() != 0:
        print(search_5digits_number())
    else:
        print('No number matched...')


def search_6digits_number():
    top_digit = [9 - i for i in range(9)]
    other_digit = [9 - i for i in range(10)]

    for top in top_digit:
        for second in other_digit:
            for third in other_digit:
                target = top * 100001 + second * 10010 + third * 1100
                result = search_3digits_divider(target)

                if result != 0:
                    print(result)
                    return search_3digits_divider(target)

    return 0


def search_5digits_number():
    top_digit = [9 - i for i in range(9)]
    other_digit = [9 - i for i in range(10)]

    for top in top_digit:
        for second in other_digit:
            for third in other_digit:
                target = top * 10001 + second * 1010 + third * 100
                result = search_3digits_divider(target)

                if result != 0:
                    print(result)
                    return search_3digits_divider(target)

    return 0


def search_3digits_divider(num):
    maximum = int(sqrt(num))

    for i in range(100, maximum + 1):
        if num % i == 0 and (num / i) >= 100 and (num / i) <= maximum:
            print(i, "times", (num / i), "equals", num)
            return num

    return 0


if __name__ == '__main__':
    main()
