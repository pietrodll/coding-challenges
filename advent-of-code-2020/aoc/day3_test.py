from .day3 import count_trees, check_slopes, parse_input

data = """..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#"""


def test_parse_input():
    expected = [
        [0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0],
        [1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0],
        [0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0],
        [0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1],
        [0, 1, 0, 0, 0, 1, 1, 0, 0, 1, 0],
        [0, 0, 1, 0, 1, 1, 0, 0, 0, 0, 0],
        [0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 1],
        [0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1],
        [1, 0, 1, 1, 0, 0, 0, 1, 0, 0, 0],
        [1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 1],
        [0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 1],
    ]

    assert parse_input(data) == expected


def test_count_trees():
    assert count_trees(parse_input(data), (3, 1)) == 7


def test_check_slopes():
    grid = parse_input(data)
    slopes = [(1, 1), (3, 1), (5, 1), (7, 1), (1, 2)]

    assert check_slopes(grid, slopes) == 336
