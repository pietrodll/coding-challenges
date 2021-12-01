"""
Problem 23
==========

A perfect number is a number for which the sum of its proper divisors is exactly equal to the
number. For example, the sum of the proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which
means that 28 is a perfect number.

A number n is called deficient if the sum of its proper divisors is less than n and it is called
abundant if this sum exceeds n.

As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number that can be
written as the sum of two abundant numbers is 24. By mathematical analysis, it can be shown that
all integers greater than 28123 can be written as the sum of two abundant numbers. However, this
upper limit cannot be reduced any further by analysis even though it is known that the greatest
number that cannot be expressed as the sum of two abundant numbers is less than this limit.

Find the sum of all the positive integers which cannot be written as the sum of two abundant
numbers.
"""

from utils.divisors import divisor_sum


def is_abundant(n):
    return divisor_sum(n) > n


def find_abundant(upper_limit):
    return list(filter(is_abundant, range(12, upper_limit)))


def find_abundant_sums(upper_limit):
    abundant = find_abundant(upper_limit)

    sums = set()

    for x in abundant:
        for y in abundant:
            s = x + y

            if s < upper_limit:
                sums.add(s)

    return sums


def find_non_abundant_sums():
    sums = find_abundant_sums(28123)

    return set(range(1, 28123)) - sums


if __name__ == "__main__":
    assert is_abundant(12)
    assert 12 in find_abundant(15)

    print(sum(find_non_abundant_sums()))
