"""
Problem 120

Let r be the remainder when (a−1)^n + (a+1)^n is divided by a^2.

For example, if a = 7 and n = 3, then r = 42: 63 + 83 = 728 ≡ 42 mod 49. And as n varies,
so too will r, but for a = 7 it turns out that rmax = 42.

For 3 ≤ a ≤ 1000, find ∑ rmax.

=========================================================

(a-1)^n = ∑ (k comb n) a^k (-1)^(n-k) = (-1)^n + n a (-1)^(n-1) [mod a²]
(a+1)^n = ∑ (k comb n) a^k = 1 + n a [mod a²]

(a-1)^n + (a+1)^n = 1 + (-1)^n + n a (1 - (-1)^n) [mod a²]
"""

def compute_r(a, n):
    if n % 2 == 0:
        return 2 % (a * a)

    return (2 * n * a) % (a * a)


def compute_rmax(a):
    rmax = 2

    for i in range(1, a * a // 2):
        r = compute_r(a, 2 * i + 1)
        # print(i, r)

        if r > rmax:
            rmax = r

    return rmax


def sum_rmax(low, up):
    tot = 0
    for a in range(low, up + 1):
        tot += compute_rmax(a)

    return tot


if __name__ == "__main__":
    print('a = 7, n = 3, r =', compute_r(7, 3))
    print(compute_rmax(7))

    print(sum_rmax(3, 1000))
