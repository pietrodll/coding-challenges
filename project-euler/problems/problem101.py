"""
Problem 101
===========
"""

import numpy as np


def compute(u, n):
    res = u[0]
    for k in range(1, len(u)):
        res += u[k] * n ** k

    return res


def generate_matrix(n, iterable=None):
    it = range(1, n + 1)

    if iterable is not None:
        it = iterable

    A = []

    for x in it:
        L = [1]

        for k in range(1, n):
            coeff = L[k - 1] * x
            L.append(coeff)

        A.append(L)

    return A


def op(u, k):
    if len(u) == 1:
        return [u[0]]

    values = [compute(u, n) for n in range(1, k + 1)]
    matrix = generate_matrix(k)

    res = np.linalg.solve(matrix, values)

    return res


def bop(u, k):
    if k >= len(u):
        return None

    approx = op(u, k)

    return compute(approx, k + 1)


def bop_sum(u):
    S = 0
    for k in range(1, len(u)):
        S += bop(u, k)
    return S


def main():
    # example = [0, 0, 0, 1]
    u = [1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1]

    print(bop_sum(u))
