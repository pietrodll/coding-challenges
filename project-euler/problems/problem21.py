"""
Problem 21
==========
"""

from math import sqrt


def divisor_sum(n):
    S = 1
    for d in range(2, int(sqrt(n))):
        if n % d == 0:
            S += d
            S += n // d
    return S


def divisor_sum_list(N):
    L = [0, 1]
    for n in range(2, N):
        L.append(divisor_sum(n))
    return L


def is_amicable(n):
    a = divisor_sum(n)
    b = divisor_sum(a)
    return n != a and n == b


def amicale_sum(N):
    S = 0
    for i in range(2, N):
        if is_amicable(i):
            S += i
    return S


def main():
    print(is_amicable(220))
    print(is_amicable(284))

    print(amicale_sum(10000))
