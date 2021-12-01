"""
Problem 41
==========

We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly
once. For example, 2143 is a 4-digit pandigital and is also prime.

What is the largest n-digit pandigital prime that exists?
"""

from utils.digits import digits_to_num
from utils.divisors import is_prime
from utils.permutations import iter_permutations


def main():
    largest = 0

    for n in range(2, 9 + 1):
        for perm in iter_permutations(range(1, n + 1)):
            x = digits_to_num(perm)

            if  x > largest and is_prime(x):
                largest = x

    print(largest)


if __name__ == "__main__":
    main()
