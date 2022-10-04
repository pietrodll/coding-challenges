"""
Problem 33
==========

The fraction 49/98 is a curious fraction, as an inexperienced mathematician in attempting to
simplify it may incorrectly believe that 49/98 = 4/8, which is correct, is obtained by cancelling
the 9s.

We shall consider fractions like, 30/50 = 3/5, to be trivial examples.

There are exactly four non-trivial examples of this type of fraction, less than one in value, and
containing two digits in the numerator and denominator.

If the product of these four fractions is given in its lowest common terms, find the value of the
denominator.
"""


def gcd(a, b):
    while b != 0:
        a, b = b, a % b

    return a


def main():
    den_prod = 1
    num_prod = 1
    fractions = []

    for common in range(1, 10):
        for den in range(1, 10):
            for num in range(1, den):
                if (10 * common + num) * den == num * (10 * common + den):
                    fractions.append(f"{10 * common + num}/{10 * common + den}")
                    den_prod *= 10 * common + den
                    num_prod *= 10 * common + num

                elif (10 * num + common) * den == num * (10 * common + den):
                    fractions.append(f"{10 * num + common}/{10 * common + den}")
                    den_prod *= 10 * common + den
                    num_prod *= 10 * num + common

                elif (10 * common + num) * den == num * (10 * den + common):
                    fractions.append(f"{10 * common + num}/{10 * den + common}")
                    den_prod *= 10 * den + common
                    num_prod *= 10 * common + num

                elif (10 * num + common) * den == num * (10 * den + common):
                    fractions.append(f"{10 * num + common}/{10 * den + common}")
                    den_prod *= 10 * den + common
                    num_prod *= 10 * num + common

    result = den_prod // gcd(den_prod, num_prod)

    print(result)
    print(fractions)
