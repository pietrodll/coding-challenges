"""
The prime 41, can be written as the sum of six consecutive primes:

41 = 2 + 3 + 5 + 7 + 11 + 13
This is the longest sum of consecutive primes that adds to a prime below one-hundred.

The longest sum of consecutive primes below one-thousand that adds to a prime, contains 21 terms,
and is equal to 953.

Which prime, below one-million, can be written as the sum of the most consecutive primes?
"""

from utils.divisors import is_prime, get_primes


def get_primes_bis(N):
    primes = [False] * N
    prime_list = []

    for i in range(N):
        if is_prime(i):
            primes[i] = True
            prime_list.append(i)
    return primes, prime_list


def longest_prime_sum_from_start(N, start, primes, prime_list):
    longest = 0
    index = start
    S = 0
    i = start

    while S < N and i < len(prime_list):
        if primes[S]:
            longest = S
            index = i

        S += prime_list[i]
        i += 1

    return longest, index - start


def longest_prime_sum(N):
    primes, prime_list = get_primes(N)

    max_len = 0
    max_len_sum = 0
    for i in range(N):
        S, length = longest_prime_sum_from_start(N, i, primes, prime_list)

        if length > max_len:
            max_len = length
            max_len_sum = S

    return max_len_sum, max_len


def check_primes(N):
    primes, prime_list = get_primes(N)
    primes_bis, prime_list_bis = get_primes_bis(N)

    assert len(prime_list) == len(prime_list_bis)
    assert len(primes) == len(primes_bis)

    n = len(prime_list)
    m = len(primes)

    for i in range(m):
        assert primes[i] == primes_bis[i]

    for i in range(n):
        assert prime_list[i] == prime_list_bis[i]


def main():
    check_primes(1000)
    print(longest_prime_sum(1000000))
