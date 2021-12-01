"""
Problem 28
==========

Starting with the number 1 and moving to the right in a clockwise direction a 5 by 5 spiral is
formed as follows:

21 22 23 24 25
20  7  8  9 10
19  6  1  2 11
18  5  4  3 12
17 16 15 14 13

It can be verified that the sum of the numbers on the diagonals is 101.

What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral formed in the same way?
"""

def diagonals(size):
    if size % 2 == 0:
        raise ValueError(f'The size of the spiral cannot be even! {size} was given.')

    down = [1]
    up = [1]

    for p in range(2, size, 2):
        for _ in range(2):
            down.append(up[-1] + p)
            up.append(down[-1] + p)

    return down, up


def diagonals_sum(size):
    if size % 2 == 0:
        raise ValueError(f'The size of the spiral cannot be even! {size} was given.')

    elem = 1
    S = 1

    for p in range(2, size, 2):
        for _ in range(4):
            elem += p
            S += elem

    return S


def test():
    down, up = diagonals(5)
    S = diagonals_sum(5)

    assert sum(down) + sum(up) == S + 1
    assert S == 101


if __name__ == "__main__":
    test()

    print(diagonals_sum(1001))
