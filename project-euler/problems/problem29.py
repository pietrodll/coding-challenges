"""
Problem 29
==========

Consider all integer combinations of ab for 2 ≤ a ≤ 5 and 2 ≤ b ≤ 5:

    2^2 = 4, 2^3 = 8, 2^4 = 16, 2^5 = 32
    3^2 = 9, 3^3 = 27, 3^4 = 81, 3^5 = 243
    4^2 = 16, 4^3 = 64, 4^4 = 256, 4^5 = 1024
    5^2 = 25, 5^3 = 125, 5^4 = 625, 5^5 = 3125

If they are then placed in numerical order, with any repeats removed, we gett he following
sequence of 15 distinct terms:

4, 8, 9, 16, 25, 27, 32, 64, 81, 125, 243, 256, 625, 1024, 3125

How many distinct terms are in the sequence generated by ab for 2 ≤ a ≤ 100 and 2 ≤ b ≤ 100?
"""

# The numbers will become too big to be represented in memory (100^100 > 2^64)
# Hence, we use the prime factor decomposition

from utils.divisors import prime_decomposition, get_primes


def distinct_powers(N):
    _, primes = get_primes(N + 1)

    decompositions = set()

    for a in range(2, N + 1):
        for b in range(2, N + 1):
            decomp = tuple(val * b for val in prime_decomposition(a, primes))
            decompositions.add(decomp)

    return decompositions


def main():
    assert tuple(prime_decomposition(2)) == (1,)
    assert tuple(prime_decomposition(10)) == (1, 0, 1, 0)

    assert len(distinct_powers(5)) == 15

    print(len(distinct_powers(100)))
