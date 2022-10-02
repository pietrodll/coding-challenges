"""Day 3"""

from typing import List, Tuple


def parse_input(data: str) -> List[List[int]]:
    result = []

    for line in data.split("\n"):
        result.append([int(c == "#") for c in line])

    return result


def navigate(grid: list, slope: Tuple[int, int]):
    right, down = slope
    i, j = 0, 0

    while i < len(grid):
        yield grid[i][j]

        i += down
        j = (j + right) % len(grid[0])


def count_trees(grid: List[List[int]], slope: Tuple[int, int]):
    return sum(navigate(grid, slope))


def check_slopes(grid: List[List[int]], slopes: List[Tuple[int, int]]):
    product = 1

    for slope in slopes:
        product *= count_trees(grid, slope)

    return product


def main(data: str):
    grid = parse_input(data)

    print("Number of encountered trees:", count_trees(grid, (3, 1)))

    slopes = [(1, 1), (3, 1), (5, 1), (7, 1), (1, 2)]
    print("Other slopes result:", check_slopes(grid, slopes))
