from .day5 import decode_seating


def test_decode_seating():
    assert decode_seating("FBFBBFFRLR") == (44, 5)
