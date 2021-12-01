"""
Problem 30
==========

Surprisingly there are only three numbers that can be written as the sum of fourth powers of their
digits:

    1634 = 1^4 + 6^4 + 3^4 + 4^4
    8208 = 8^4 + 2^4 + 0^4 + 8^4
    9474 = 9^4 + 4^4 + 7^4 + 4^4

As 1 = 1^4 is not a sum it is not included.

The sum of these numbers is 1634 + 8208 + 9474 = 19316.

Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.
"""

# This equality is possible only when the number has less than k digits:
# N = sum(a_i * 10^i) > a_k * 10^k, and sum(a_i^5) < sum(10^5) = k * 10^5
# Hence, if there is equality: a_k * 10^k < k * 10^5 (where 0 < a_k < 10) which can only be true for
# k <= 5


def is_digit_power_sum(N, pw):
    n = N
    S = 0

    while n != 0:
        S += (n % 10) ** pw
        n //= 10

    return S == N


def find_digit_power_sum(pw):
    sums = []

    for N in range(2, 10 ** (pw + 1)):
        if is_digit_power_sum(N, pw):
            sums.append(N)

    return sums


if __name__ == "__main__":
    assert is_digit_power_sum(1634, 4) is True
    assert is_digit_power_sum(8208, 4) is True
    assert is_digit_power_sum(9474, 4) is True

    print(sum(find_digit_power_sum(5)))
