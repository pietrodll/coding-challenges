"""
Problem 74
==========

The number 145 is well known for the property that the sum of the factorial of its digits is equal to 145:

1! + 4! + 5! = 1 + 24 + 120 = 145

Perhaps less well known is 169, in that it produces the longest chain of numbers that link back to 169; it turns out that there are only three such loops that exist:

169 → 363601 → 1454 → 169
871 → 45361 → 871
872 → 45362 → 872

It is not difficult to prove that EVERY starting number will eventually get stuck in a loop. For example,

69 → 363600 → 1454 → 169 → 363601 (→ 1454)
78 → 45360 → 871 → 45361 (→ 871)
540 → 145 (→ 145)

Starting with 69 produces a chain of five non-repeating terms, but the longest non-repeating chain with a starting number below one million is sixty terms.

How many chains, with a starting number below one million, contain exactly sixty non-repeating terms?
"""

import math
from typing import Generator, List, Set


def iter_digits(n: int) -> Generator[int, None, None]:
    while n != 0:
        yield n % 10
        n = n // 10


def compute_digit_factorial_sum(n: int) -> int:
    return sum(map(math.factorial, iter_digits(n)))


def find_loop(n: int) -> Set[int]:
    loop = {n}
    nxt = compute_digit_factorial_sum(n)

    while nxt not in loop:
        loop.add(nxt)
        nxt = compute_digit_factorial_sum(nxt)

    return loop


def find_all_loops(max_number: int) -> List[Set[int]]:
    loops = [set() for _ in range(max_number + 1)]

    for n in range(1, max_number + 1):
        loop = {n}
        nxt = compute_digit_factorial_sum(n)

        while nxt not in loop:
            if nxt < len(loops) and len(loops[nxt]) > 0:
                loop |= loops[nxt]
            else:
                loop.add(nxt)
                nxt = compute_digit_factorial_sum(nxt)

        loops[n] = loop

    return loops


def main():
    all_loops = find_all_loops(10 ** 6)
    loops_with_length_60 = []

    for i, loop in enumerate(all_loops):
        if len(loop) == 60:
            loops_with_length_60.append(i)

    print(len(loops_with_length_60))
