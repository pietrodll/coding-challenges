"""Day 10"""

from collections import Counter
from typing import List, Tuple


def parse_input(data: str) -> List[int]:
    return list(map(int, data.splitlines()))


def count_differences(adapters: List[int]) -> Tuple[int, int]:
    sequence = list(adapters)
    sequence.sort()

    counter = Counter()

    # the plug has a zero jolt
    counter[sequence[0]] += 1

    for i in range(len(sequence) - 1):
        counter[sequence[i + 1] - sequence[i]] += 1

    # do not forget the difference between the device and the last adapter
    counter[3] += 1

    return counter[1], counter[3]


def count_arrangements(adapters: List[int]) -> int:
    sequence = list(adapters)
    sequence.sort()

    def aux(seq: List[int], i: int, memo: List[int]) -> int:
        if i >= len(seq):
            return 0

        if i == len(seq) - 1:
            return 1

        if memo[i] is not None:
            return memo[i]

        # find the next adapter with a difference strictly greater than 3
        j = i + 1

        while j < len(seq) and seq[j] - seq[i] <= 3:
            j += 1

        if j == i + 1:
            raise ValueError(
                "Difference between consecutive values cannot be more than three"
            )

        result = 0

        for k in range(i + 1, j):
            # we can skip any adapter between i and j excluded
            result += aux(seq, k, memo)

        memo[i] = result
        return result

    return aux([0] + sequence, 0, [None] * (len(sequence) + 1))


def main(data: str):
    adapters = parse_input(data)

    diff1, diff3 = count_differences(adapters)

    print("Differences:", diff1 * diff3)
    print("Number of arrangements:", count_arrangements(adapters))
