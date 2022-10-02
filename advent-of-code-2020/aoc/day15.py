"""Day 15"""

from typing import List


def parse_input(data: str):
    return list(map(int, data.split(",")))


def play_game(starting_numbers: List[int], stop: int) -> int:
    last_indexes = {}
    last_spoken = starting_numbers[0]
    turn = 1

    while turn < len(starting_numbers):
        last_indexes[last_spoken] = turn - 1
        last_spoken = starting_numbers[turn]
        turn += 1

    while turn < stop:
        if last_spoken in last_indexes:
            number = turn - 1 - last_indexes[last_spoken]
        else:
            number = 0

        last_indexes[last_spoken] = turn - 1
        last_spoken = number
        turn += 1

    return last_spoken


def main(data: str):
    starting_numbers = parse_input(data)

    print("2020th number:", play_game(starting_numbers, 2020))
    print("30000000th number:", play_game(starting_numbers, 30000000))
