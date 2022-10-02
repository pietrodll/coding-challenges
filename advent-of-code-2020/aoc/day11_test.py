from .day11 import (
    SeatSimulator,
    SeatSimulatorV2,
    get_visible_seats,
    parse_input,
)

data = """L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL"""

expected_next_states = [
    """#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##""",
    """#.LL.L#.##
#LLLLLL.L#
L.L.L..L..
#LLL.LL.L#
#.LL.LL.LL
#.LLLL#.##
..L.L.....
#LLLLLLLL#
#.LLLLLL.L
#.#LLLL.##""",
    """#.##.L#.##
#L###LL.L#
L.#.#..#..
#L##.##.L#
#.##.LL.LL
#.###L#.##
..#.#.....
#L######L#
#.LL###L.L
#.#L###.##""",
    """#.#L.L#.##
#LLL#LL.L#
L.L.L..#..
#LLL.##.L#
#.LL.LL.LL
#.LL#L#.##
..L.L.....
#L#LLLL#L#
#.LLLLLL.L
#.#L#L#.##""",
    """#.#L.L#.##
#LLL#LL.L#
L.#.L..#..
#L##.##.L#
#.#L.LL.LL
#.#L#L#.##
..L.L.....
#L#L##L#L#
#.LLLLLL.L
#.#L#L#.##""",
]


def test_next_state_adjacent_seats():
    sim = SeatSimulator(parse_input(data))

    for expected in map(parse_input, expected_next_states):
        sim.next_state()
        assert sim.grid == expected


def test_count_occupied_seats_at_equilibrium_adjacent_seats():
    sim = SeatSimulator(parse_input(data))
    sim.compute_until_equilibrium()

    assert sim.count_occupied_seats() == 37


expected_next_states_visible_seats = [
    """#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##""",
    """#.LL.LL.L#
#LLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLLL.L
#.LLLLL.L#""",
    """#.L#.##.L#
#L#####.LL
L.#.#..#..
##L#.##.##
#.##.#L.##
#.#####.#L
..#.#.....
LLL####LL#
#.L#####.L
#.L####.L#""",
    """#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##LL.LL.L#
L.LL.LL.L#
#.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLL#.L
#.L#LL#.L#""",
    """#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##L#.#L.L#
L.L#.#L.L#
#.L####.LL
..#.#.....
LLL###LLL#
#.LLLLL#.L
#.L#LL#.L#""",
    """#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##L#.#L.L#
L.L#.LL.L#
#.LLLL#.LL
..#.L.....
LLL###LLL#
#.LLLLL#.L
#.L#LL#.L#""",
]


def test_visible_seats_iterator():
    grid = parse_input(
        ".......#.\n"
        "...#.....\n"
        ".#.......\n"
        ".........\n"
        "..#L....#\n"
        "....#....\n"
        ".........\n"
        "#........\n"
        "...#....."
    )
    visible_seats = get_visible_seats(4, 3, grid)

    assert len(list(visible_seats)) == 8

    grid = parse_input(
        ".##.##.\n"
        "#.#.#.#\n"
        "##...##\n"
        "...L...\n"
        "##...##\n"
        "#.#.#.#\n"
        ".##.##."
    )
    visible_seats = get_visible_seats(3, 3, grid)

    assert list(visible_seats) == []

    grid = parse_input(".............\n.L.L.#.#.#.#.\n.............")
    visible_seats = get_visible_seats(1, 1, grid)

    assert list(visible_seats) == [(1, 3)]

    grid = parse_input(
        "#.##.##.##\n"
        "#######.##\n"
        "#.#.#..#..\n"
        "####.##.##\n"
        "#.##.##.##\n"
        "#.#####.##\n"
        "..#.#.....\n"
        "##########\n"
        "#.######.#\n"
        "#.#####.##"
    )
    visible_seats = get_visible_seats(0, 2, grid)
    assert set(visible_seats) == {(0, 0), (0, 3), (1, 1), (1, 2), (1, 3)}


def test_next_state_visible_seats():
    sim = SeatSimulatorV2(parse_input(data))

    for expected in map(parse_input, expected_next_states_visible_seats):
        sim.next_state()
        assert sim.grid == expected
