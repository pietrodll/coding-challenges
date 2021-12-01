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

def hcf(a, b):
    """Computes the highest common factor of two numbers"""

    if a < b:
        a, b = b, a

    while b > 0:
        r = a % b
        a = b
        b = r
    return a


def count_fractions(N):
    """Computes the number of reduced proper fractions with a denominator less or equal than N"""

    num = 0
    for d in range(1, N+1):
        print(d, end='\r')
        for n in range(1, d):
            if hcf(n, d) == 1:
                num += 1
    return num


def test():
    print('hcf(10, 8) =', hcf(10, 8))
    print('hcf(25, 15) =', hcf(25, 15))
    print('hcf(25, 12) =', hcf(25, 12))
    print('count_fractions(8) =', count_fractions(8))


def main():
    print(count_fractions(10000))
