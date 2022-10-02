from .day1 import parse_input, find_three_entries_summing_to, find_entries_summing_to

data = """1721
979
366
299
675
1456"""


def test_parse_input():
    expected = [1721, 979, 366, 299, 675, 1456]
    assert parse_input(data) == expected


def test_find_entries():
    assert find_entries_summing_to(parse_input(data), 2020) == (299, 1721)


def test_find_three_entries():
    assert find_three_entries_summing_to(parse_input(data), 2020) == (675, 366, 979)
