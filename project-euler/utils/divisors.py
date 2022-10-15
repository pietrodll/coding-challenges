"""This module contains utility functions for divisors and prime numbers"""

from math import sqrt
from typing import List, Tuple


def is_prime(n):
    for i in range(2, int(sqrt(n)) + 1):
        if n % i == 0:
            return False

    return True


def get_primes(N: int) -> Tuple[List[bool], List[int]]:
    primes = [False] * 2 + [True] * (N - 2)

    i = 2
    while i < N:
        j = 2

        while j * i < N:
            primes[j * i] = False
            j += 1

        i += 1

        while i < N and not primes[i]:
            i += 1

    prime_list = [i for i in range(N) if primes[i]]

    return primes, prime_list


def divisors(n, sort=False):
    div = [1]
    sqr = int(sqrt(n))

    for i in range(2, sqr + 1):
        if n % i == 0:
            div.append(i)

            if i != sqr:
                div.append(n // i)

    if sort:
        div.sort()

    return div


def divisor_sum(n):
    S = 1
    sqr = sqrt(n)

    for i in range(2, int(sqr) + 1):
        if n % i == 0:
            S += i

            if i != sqr:
                S += n // i

    return S


def mult_ord(a, n):
    """Returns the multiplicative order of a modulo n. Raises ValueError if a is not relatively
    prime to n
    """

    if n == 1:
        raise ValueError("The modulo cannot be equal to 1")

    m_ord = 1
    value = a % n

    while value != 1 and m_ord < n:
        value = (value * a) % n
        m_ord += 1

    if m_ord < n:
        return m_ord

    raise ValueError(
        f"The given integer arguments are not relative primes: {a} and {n}"
    )


def prime_decomposition(N, primes=None):
    if primes is None:
        _, primes = get_primes(N + 1)

    decomp = []

    for p in primes:
        val = 0

        while N % p == 0:
            N //= p
            val += 1

        decomp.append(val)

    return decomp
