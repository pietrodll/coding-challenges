"""
Problem 36
==========


The decimal number, 585 = 1001001001 (binary), is palindromic in both bases.

Find the sum of all numbers, less than one million, which are palindromic in base 10 and base 2.

(Please note that the palindromic number, in either base, may not include leading zeros)
"""


def get_digits(n, base):
    digits = []

    while n > 0:
        digits.append(n % base)
        n //= base

    return digits


def is_palindromic(L):
    return all(L[i] == L[-1 - i] for i in range(len(L) // 2))


def is_double_palindromic(n):
    return is_palindromic(get_digits(n, 10)) and is_palindromic(get_digits(n, 2))


def test():
    assert is_double_palindromic(585)


def main():
    S = 0

    for n in range(1, 1000000):
        if is_double_palindromic(n):
            S += n

    print(S)


if __name__ == "__main__":
    test()
    main()
