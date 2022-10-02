"""Day 9"""

from itertools import islice
from typing import List


def parse_input(data: str) -> List[int]:
    return list(map(int, data.splitlines()))


def is_sum(value: int, numbers: List[int]):
    numbers_set = set(numbers)

    for x in numbers_set:
        compl = value - x

        if x != compl and compl in numbers_set:
            return True

    return False


def find_first_invalid(sequence: List[int], preamble_size: int):
    for i in range(preamble_size, len(sequence)):
        val = sequence[i]

        if not is_sum(val, islice(sequence, i - preamble_size, i)):
            return val

    return None


def find_contiguous_subset_summing_to(sequence: List[int], target: int):
    for start in range(len(sequence)):
        tot = 0
        i = start

        while tot < target and i < len(sequence):
            tot += sequence[i]
            i += 1

        if tot == target:
            return [sequence[k] for k in range(start, i)]

    raise ValueError(f"Cannot find contiguous subset summing to {target}")


def find_encryption_weakness(sequence: List[int], first_invalid: int):
    subset = find_contiguous_subset_summing_to(sequence, first_invalid)

    return max(subset) + min(subset)


def main(data: str):
    sequence = parse_input(data)

    first_invalid = find_first_invalid(sequence, 25)

    print("First invalid number:", first_invalid)

    print("Encryption weakness:", find_encryption_weakness(sequence, first_invalid))
