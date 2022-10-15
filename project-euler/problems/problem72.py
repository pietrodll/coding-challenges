"""
Problem 72
==========

Consider the fraction, n/d, where n and d are positive integers.
If n<d and HCF(n,d)=1, it is called a reduced proper fraction.

If we list the set of reduced proper fractions for d ≤ 8 in ascending order of size, we get:

1/8, 1/7, 1/6, 1/5, 1/4, 2/7, 1/3, 3/8, 2/5, 3/7, 1/2,
4/7, 3/5, 5/8, 2/3, 5/7, 3/4, 4/5, 5/6, 6/7, 7/8

It can be seen that there are 21 elements in this set.

How many elements would be contained in the set of reduced proper fractions for d ≤ 1,000,000?
"""

# The number of reduced fractions with denominator d is the number of integers up to d that are
# relatively prime to d. This is given by Euler's phi function.
# Therefore, the result is the sum of the values of the phi function from 2 to 1,000,000

from typing import List


def euler_phi_values(N: int) -> List[int]:
    phi = list(range(N + 1))

    for i in range(2, N + 1):
        if phi[i] == i:
            for j in range(i, N + 1, i):
                phi[j] = (phi[j] // i) * (i - 1)

    return phi


def count_fractions(N: int) -> int:
    """Computes the number of reduced proper fractions with a denominator less or equal than N"""

    # We exclude phi[1] = 1
    return sum(euler_phi_values(N)) - 1


def main():
    print(count_fractions(1_000_000))
