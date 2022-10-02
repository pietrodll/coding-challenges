from .day6 import count_group_all_answers, parse_input, count_group_different_answers

data = """abc

a
b
c

ab
ac

a
a
a
a

b"""


def test_parse_input():
    expected = [["abc"], ["a", "b", "c"], ["ab", "ac"], ["a", "a", "a", "a"], ["b"]]

    assert parse_input(data) == expected


def test_count_different_answers():
    answers = parse_input(data)
    expected_counts = [3, 3, 3, 1, 1]

    for answer, expected_count in zip(answers, expected_counts):
        assert count_group_different_answers(answer) == expected_count


def test_all_answered_count():
    answers = parse_input(data)
    expected_counts = [3, 0, 1, 1, 1]

    for answer, expected_count in zip(answers, expected_counts):
        assert count_group_all_answers(answer) == expected_count
