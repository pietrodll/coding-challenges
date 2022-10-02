from .day15 import play_game


def test_play_game():
    assert play_game([0, 3, 6], 10) == 0
    assert play_game([0, 3, 6], 2020) == 436
    assert play_game([1, 3, 2], 2020) == 1
    assert play_game([2, 1, 3], 2020) == 10
    assert play_game([1, 2, 3], 2020) == 27
    assert play_game([2, 3, 1], 2020) == 78
    assert play_game([3, 2, 1], 2020) == 438
    assert play_game([3, 1, 2], 2020) == 1836
