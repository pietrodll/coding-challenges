"""
Euler discovered the remarkable quadratic formula: n^2 + n + 41

It turns out that the formula will produce 40 primes for the consecutive integer values
0 <= n <= 39. However, when n = 40, 40^2 + 40 + 41 = 40 * (40 + 1) + 41 is divisible by 41, and
certainly when n = 41, 41^2 + 41 + 41 is clearly divisible by 41.

The incredible formula n^2 - 79 * n + 1601 was discovered, which produces 80 primes for the
consecutive values 0 <= n <= 79. The product of the coefficients, −79 and 1601, is −126479.

Considering quadratics of the form: n^2 + a * n + b, where abs(a) < 1000 and abs(b) < 1000

Find the product of the coefficients, a and b, for the quadratic expression that produces the
maximum number of primes for consecutive values of n, starting with n = 0.
"""

from utils.divisors import get_primes


# As the formula must give a prime number for n = 0, it means that b must be a prime number
# If n = b, the result is divisible by b, which means that n varies between 0 and b - 1
# Depending on the value of a, the maximum value of the formula can be b (with n = 0)
# or (b - 1)^2 + a * (b - 1) + b = b^2 + ab - a - b + 1 = b^2 + (a - 1) * (b - 1) (with n = b - 1),
# which means that we only have to compute the prime numbers lower than
# max(b, b^2 + (a - 1) * (b - 1)) = bmax^2 + (amax - 1) * (bmax - 1) to solve the problem


def value(a, b, n):
    return n * n + a * n + b


def count_produced_primes(a, b, primes):
    count = 0

    for n in range(b):
        val = value(a, b, n)

        if val < 0 or not primes[val]:
            return count

        count += 1

    return count


def find_best_coefficients(amax, bmax):
    primes, prime_list = get_primes(bmax * bmax + (amax - 1) * (bmax - 1))

    abest = 0
    bbest = prime_list[0]
    maxcount = 0

    b = 0
    while b < bmax and b < len(prime_list):
        for a in range(-amax + 1, amax):
            count = count_produced_primes(a, b, primes)

            if count > maxcount:
                maxcount = count
                abest = a
                bbest = b

        b += 1

    return abest, bbest


def find_best_coeff_product(amax, bmax):
    a, b = find_best_coefficients(amax, bmax)

    return a * b


def main():
    print(find_best_coeff_product(1000, 1000))
