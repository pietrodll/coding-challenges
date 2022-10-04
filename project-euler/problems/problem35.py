"""
Problem 35
==========

The number, 197, is called a circular prime because all rotations of the digits: 197, 971, and 719, are themselves prime.

There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.

How many circular primes are there below one million?
"""

from typing import List
from utils.digits import get_digits, digits_to_num
from utils.divisors import get_primes


def rotate_digits(n: int) -> List[int]:
    digits = get_digits(n)
    size = len(digits)
    rotations = []

    for shift in range(size):
        rotated_digits = [digits[(i + shift) % size] for i in range(size)]
        rotations.append(digits_to_num(rotated_digits))

    return rotations


def count_circular_primes_below(n: int):
    primes, primes_list = get_primes(n + 1)
    count = 0

    for prime in primes_list:
        if all(primes[rotated] for rotated in rotate_digits(prime)):
            count += 1

    return count


def main():
    print(count_circular_primes_below(1000000))
