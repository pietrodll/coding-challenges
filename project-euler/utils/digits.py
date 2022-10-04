"""This module contains utility functions for digits"""


from typing import Generator, List


def get_digits(n: int, reverse: bool = False) -> List[int]:
    """Returns the digits of n as a list, from the lowest power of 10 to the highest. If `reverse`
    is `True`, returns them in from the highest to the lowest"""

    if reverse:
        return list(reversed(iter_digits(n)))

    return list(iter_digits(n))


def iter_digits(n: int) -> Generator[int, None, None]:
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


def digits_to_num(L: List[int], reverse: bool = False) -> int:
    """Returns a number from a list of digits, given by the lowest power of 10 to the highest, or
    the other way around if `reverse` is True"""

    digits = reversed(L) if reverse else L
    n = 0
    pow_of_10 = 1

    for d in digits:
        n += d * pow_of_10
        pow_of_10 *= 10

    return n
