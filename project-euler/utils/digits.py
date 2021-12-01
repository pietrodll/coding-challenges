"""This module contains utility functions for digits"""

def get_digits(n, reverse=False):
    """Returns the digits of n as a list, from the lowest power of 10 to the highest. If `reverse`
    is `True`, returns them in from the highest to the lowest"""

    digits = []

    while n != 0:
        digits.append(n % 10)
        n //= 10

    if reverse:
        digits.reverse()

    return digits


def iter_digits(n):
    """Iterates through the the digits of n, from the lowest power of 10 to the highest"""

    while n != 0:
        yield n % 10

        n //= 10


def is_pandigital_seq(sequence):
    digits = set(range(1, 10))

    for x in sequence:
        if x not in digits:
            return False

        digits.remove(x)

    return len(digits) == 0


def is_pandigital(n):
    return is_pandigital_seq(iter_digits(n))


def digits_to_num(L, reverse=False):
    """Returns a number from a list of digits, given by the lowest power of 10 to the highest, or
    the other way around if `reverse` is True"""

    digits = reversed(L) if reverse else L
    n = 0

    for i, d in enumerate(digits):
        n += d * (10 ** i)

    return n
