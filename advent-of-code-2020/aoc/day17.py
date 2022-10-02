"""Day 17"""

from collections import namedtuple
from typing import Iterable, Set, Tuple

Point = namedtuple("Point", "x y z")


def parse_input(data: str) -> Set[Point]:
    points = set()

    for x, line in enumerate(data.splitlines()):
        for y, point_str in enumerate(line):
            if point_str == "#":
                points.add(Point(x, y, 0))

    return points


def find_points_ranges(
    points: Iterable[Point],
) -> Tuple[Tuple[int, int], Tuple[int, int], Tuple[int, int]]:
    for point in points:
        mins = [point[index] for index in range(len(point))]
        maxs = [point[index] for index in range(len(point))]
        break

    for point in points:
        for index in range(len(point)):
            if point[index] < mins[index]:
                mins[index] = point[index]

            if point[index] > maxs[index]:
                maxs[index] = point[index]

    return tuple(zip(mins, maxs))


def get_neighbors(point: Point) -> Set[Point]:
    neighbors = set()

    for index, value in enumerate(point):
        pass

    return neighbors


def main(data: str):
    pass
