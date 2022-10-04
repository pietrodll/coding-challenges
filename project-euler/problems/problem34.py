"""
Problem 34
==========

145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.

Find the sum of all numbers which are equal to the sum of the factorial of their digits.

Note: As 1! = 1 and 2! = 2 are not sums they are not included.
"""

# The sum of the factorials of the digits of a number N is less than 9! * log(N)
# We're looking for a number such that N = N1 N2 N3 ... Nk = N1! + N2! + N3! + ... + Nk!
# Therefore N < 9! * k => there is a maximum number above which it's useless to search (because
# the log grows more slowly than a linear function)

import math
from typing import Generator


def find_max_number_to_check() -> int:
    n = 3
    fact9 = math.factorial(9)

    while n < fact9 * math.log10(n):
        n += 1

    return n


def iter_digits(n: int) -> Generator[int, None, None]:
    while n != 0:
        yield n % 10
        n = n // 10


def compute_digit_factorial_sum(n: int) -> int:
    S = 0

    for d in iter_digits(n):
        S += math.factorial(d)

    return S


def find_result():
    M = find_max_number_to_check()
    res = 0

    for n in range(3, M + 1):
        if n == compute_digit_factorial_sum(n):
            res += n

    return res


def main():
    print(find_result())
