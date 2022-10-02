from .day13 import (
    bezout,
    find_contest_timestamp,
    find_earliest_possible_departure,
    parse_input,
    solution_two_moduli,
)


def test_parse_input():
    timestamp, schedule = parse_input("939\n7,13,x,x,59,x,31,19")
    assert timestamp == 939
    assert schedule == [(0, 7), (1, 13), (4, 59), (6, 31), (7, 19)]


def test_find_earliest_possible_departure():
    bus_id, earliest = find_earliest_possible_departure(939, [7, 13, 59, 31, 19])
    assert earliest == 944
    assert (earliest - 939) * bus_id == 295


def test_bezout():
    assert bezout(3, 4) == (-1, 1)
    assert bezout(4, 3) == (1, -1)
    assert bezout(17, 13) == (-3, 4)


def test_solution_two_moduli():
    assert solution_two_moduli(0, 3, 3, 4) == (3, 12)


def test_find_contest_timestamp():
    assert find_contest_timestamp([(0, 17), (2, 13), (3, 19)]) == 3417
    assert find_contest_timestamp([(0, 67), (1, 7), (2, 59), (3, 61)]) == 754018
    assert find_contest_timestamp([(0, 67), (2, 7), (3, 59), (4, 61)]) == 779210
    assert find_contest_timestamp([(0, 67), (1, 7), (3, 59), (4, 61)]) == 1261476
    assert (
        find_contest_timestamp([(0, 1789), (1, 37), (2, 47), (3, 1889)]) == 1202161486
    )
