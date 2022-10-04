"""
Problem 24
==========

A permutation is an ordered arrangement of objects. For example, 3124 is one possible permutation
of the digits 1, 2, 3 and 4. If all of the permutations are listed numerically or alphabetically, we
call it lexicographic order. The lexicographic permutations of 0, 1 and 2 are:

012   021   102   120   201   210

What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
"""

from math import factorial
from utils.permutations import permutations


def digit_permutations():
    M = permutations(10)
    return M


def digits_to_number(digits):
    n = 0

    for i in range(len(digits)):
        n += digits[-1 - i] * 10 ** i

    return n


def ordered_permutations():
    perms = list(map(digits_to_number, digit_permutations()))

    perms.sort()

    return perms


def main():
    assert len(permutations(5)) == factorial(5)
    assert digits_to_number([1, 2, 3]) == 123
    assert digits_to_number([0, 3, 2]) == 32
    assert digits_to_number([1, 2, 0]) == 120

    ordered = ordered_permutations()

    assert len(ordered) == factorial(10)

    print(ordered[1000000 - 1])
