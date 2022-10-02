from .day2 import check_password, check_password_v2, parse_input, Policy

data = """1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc"""


def test_parse_input():
    expected = [
        (Policy("a", 1, 3), "abcde"),
        (Policy("b", 1, 3), "cdefg"),
        (Policy("c", 2, 9), "ccccccccc"),
    ]

    assert parse_input(data) == expected


def test_check_password():
    assert check_password(Policy("a", 1, 3), "abcde") is True
    assert check_password(Policy("b", 1, 3), "cdefg") is False
    assert check_password(Policy("c", 2, 9), "ccccccccc") is True


def test_check_password_v2():
    assert check_password_v2(Policy("a", 1, 3), "abcde") is True
    assert check_password_v2(Policy("b", 1, 3), "cdefg") is False
    assert check_password_v2(Policy("c", 2, 9), "ccccccccc") is False
