"""
Problem 40
==========

An irrational decimal fraction is created by concatenating the positive integers:

0.123456789101112131415161718192021...

It can be seen that the 12th digit of the fractional part is 1.

If dn represents the nth digit of the fractional part, find the value of the following expression.

d1 × d10 × d100 × d1000 × d10000 × d100000 × d1000000
"""


def iter_digits(n):
    """Gives the digits of n as an iterator"""

    digits = []

    while n != 0:
        digits.append(n % 10)
        n //= 10

    return reversed(digits)


def digit_sequence(limit):
    """Iterator for the sequence of the concatenated digits of the natural numbers"""

    i = 0
    n = 1
    while i < limit:
        for digit in iter_digits(n):
            yield digit
            i += 1

            if i >= limit:
                break

        n += 1


def mutiply_values(iterable, indexes):
    """Mutliplies the values of the iterable given by a list of indexes"""

    indexes = set(indexes)
    product = 1

    for i, x in enumerate(iterable):
        if i in indexes:
            product *= x
            indexes.remove(i)

        if len(indexes) == 0:
            break

    return product


def main():
    indexes = [10 ** i - 1 for i in range(7)]

    result = mutiply_values(digit_sequence(1000000), indexes)

    print(result)
