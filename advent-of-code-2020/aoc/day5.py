"""Day 5"""


from typing import List, Tuple


def parse_input(data: str) -> List[int]:
    return data.splitlines()


def decode_seating(encoded: str) -> Tuple[int, int]:
    max_row = 127
    min_row = 0

    for letter in encoded[:7]:
        mid = (min_row + max_row) // 2

        if letter == "F":
            max_row = mid

        else:
            min_row = mid

    max_col = 7
    min_col = 0

    for letter in encoded[7:]:
        mid = (min_col + max_col) // 2

        if letter == "L":
            max_col = mid

        else:
            min_col = mid

    return max_row, max_col


def compute_seating_id(encoded: str) -> int:
    row, col = decode_seating(encoded)
    return row * 8 + col


def find_missing_seating(seating_ids: List[int]) -> int:
    seating_ids = list(seating_ids)
    seating_ids.sort()

    for i in range(len(seating_ids) - 1):
        if seating_ids[i + 1] - seating_ids[i] > 1:
            return seating_ids[i] + 1

    raise ValueError("Cannot find seating")


def main(data: str):
    encoded_seatings = parse_input(data)
    seating_ids = list(map(compute_seating_id, encoded_seatings))

    max_seating_id = max(seating_ids)
    missing_seating = find_missing_seating(seating_ids)

    print("Highest seating ID:", max_seating_id)
    print("Missing seating:", missing_seating)
