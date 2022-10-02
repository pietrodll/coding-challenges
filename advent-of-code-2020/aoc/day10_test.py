from .day10 import count_arrangements, count_differences

adapters = [16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4]
adapters_2 = [
    28,
    33,
    18,
    42,
    31,
    14,
    46,
    20,
    48,
    47,
    24,
    23,
    49,
    45,
    19,
    38,
    39,
    11,
    1,
    32,
    25,
    35,
    8,
    17,
    7,
    9,
    4,
    2,
    34,
    10,
    3,
]


def test_count_differences():
    assert count_differences(adapters) == (7, 5)
    assert count_differences(adapters_2) == (22, 10)


def test_count_arrangements():
    assert count_arrangements(adapters) == 8
    assert count_arrangements(adapters_2) == 19208
