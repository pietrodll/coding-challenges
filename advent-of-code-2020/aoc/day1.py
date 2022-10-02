"""Day 1"""

from typing import List


def parse_input(data: str) -> List[int]:
    return list(map(int, data.splitlines()))


def find_entries_summing_to(numbers: List[int], value: int):
    n = len(numbers)

    for i in range(n):
        target = value - numbers[i]

        for j in range(i):
            if numbers[j] == target:
                return numbers[i], numbers[j]

    raise ValueError(f"Cannot find two entries that add up to {value}")


def find_three_entries_summing_to(numbers: List[int], value: int):
    n = len(numbers)

    for i in range(n):
        for j in range(i):
            for k in range(j):
                if numbers[i] + numbers[j] + numbers[k] == value:
                    return numbers[i], numbers[j], numbers[k]

    raise ValueError(f"Cannot find three entries that add up to {value}")


def main(data: str):
    numbers = parse_input(data)

    x, y = find_entries_summing_to(numbers, 2020)
    print("Result:", x * y)

    x, y, z = find_three_entries_summing_to(numbers, 2020)
    print("Result by three:", x * y * z)
