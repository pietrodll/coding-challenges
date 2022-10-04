"""
Problem 112
===========

Working from left-to-right if no digit is exceeded by the digit to its left it is called an
increasing number; for example, 134468.

Similarly if no digit is exceeded by the digit to its right it is called a decreasing number;
for example, 66420.

We shall call a positive integer that is neither increasing nor decreasing a "bouncy" number;
for example, 155349.

As n increases, the proportion of bouncy numbers below n increases such that there are only 12951
numbers below one-million that are not bouncy and only 277032 non-bouncy numbers below 10^10.

How many numbers below a googol (10^100) are not bouncy?
"""


def binom(n, k):
    if k == 1:
        return n
    if k == 0:
        return 1
    return n * binom(n - 1, k - 1) // k


def count_increasing(n_digit):
    return binom(n_digit + 9, 9) - 1


def count_decreasing(n_digit):
    return binom(n_digit + 10, 10) - n_digit - 1


def count_non_bouncy(n_digit):
    return count_increasing(n_digit) + count_decreasing(n_digit) - 9 * n_digit


def run():
    result = count_non_bouncy(100)
    print(result)
