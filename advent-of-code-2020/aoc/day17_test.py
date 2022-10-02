from .day17 import Point, find_points_ranges


def test_find_points_ranges():
    points = {Point(0, 0, 0), Point(0, 2, 0), Point(3, 0, 0)}
    assert find_points_ranges(points) == ((0, 3), (0, 2), (0, 0))
