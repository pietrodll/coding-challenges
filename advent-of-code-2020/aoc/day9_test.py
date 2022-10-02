from .day9 import (
    find_contiguous_subset_summing_to,
    find_encryption_weakness,
    is_sum,
    find_first_invalid,
)


def test_is_sum():
    numbers = list(range(1, 26))

    assert is_sum(26, numbers) is True
    assert is_sum(40, numbers) is True
    assert is_sum(100, numbers) is False
    assert is_sum(50, numbers) is False


sequence = [
    35,
    20,
    15,
    25,
    47,
    40,
    62,
    55,
    65,
    95,
    102,
    117,
    150,
    182,
    127,
    219,
    299,
    277,
    309,
    576,
]


def test_find_first_invalid():
    assert find_first_invalid(sequence, 5) == 127


def test_find_subset():
    assert find_contiguous_subset_summing_to(sequence, 127) == [15, 25, 47, 40]


def test_find_encryption_weakness():
    assert find_encryption_weakness(sequence, 127) == 62
