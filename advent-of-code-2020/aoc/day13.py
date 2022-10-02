"""Day 13"""

from typing import List, Tuple
from math import ceil, gcd


def parse_input(data: str) -> Tuple[int, List[Tuple[int, int]]]:
    split = data.splitlines()
    schedule = [
        (offset, int(bus_id))
        for offset, bus_id in enumerate(split[1].split(","))
        if bus_id != "x"
    ]

    return int(split[0]), schedule


def find_earliest_possible_departure(
    timestamp: int, bus_ids: List[int]
) -> Tuple[int, int]:
    earliest_possible_departures = []

    for bus_id in bus_ids:
        earliest_possible_departures.append(int(ceil(timestamp / bus_id) * bus_id))

    earliest = earliest_possible_departures[0]
    best_bus = bus_ids[0]

    for i, departure in enumerate(earliest_possible_departures):
        if departure < earliest:
            earliest = departure
            best_bus = bus_ids[i]

    return best_bus, earliest


def find_waiting_time_multiplied_by_bus_id(
    timestamp: int, schedule: List[Tuple[int, int]]
) -> int:
    bus_id, earliest = find_earliest_possible_departure(
        timestamp, [bus_id for _, bus_id in schedule]
    )

    return (earliest - timestamp) * bus_id


def are_numbers_relative_primes(numbers: List[int]) -> bool:
    for i in range(len(numbers) - 1):
        for j in range(i + 1, len(numbers)):
            if gcd(numbers[i], numbers[j]) != 1:
                return False

    return True


def is_valid_contest_timestamp(schedule: List[Tuple[int, int]], timestamp: int) -> bool:
    return all((timestamp + offset) % bus_id == 0 for offset, bus_id in schedule)


def bezout(a: int, b: int) -> Tuple[int, int]:
    u, v, u1, v1 = 1, 0, 0, 1

    while b != 0:
        q = a // b
        a, u, v, b, u1, v1 = b, u1, v1, a - q * b, u - q * u1, v - q * v1

    if a != 1:
        raise ValueError("Numbers were not relatively prime")

    return u, v


def solution_two_moduli(a1: int, b1: int, a2: int, b2: int) -> Tuple[int, int]:
    m1, m2 = bezout(b1, b2)
    sol = a2 * m1 * b1 + a1 * m2 * b2
    prod = b1 * b2
    return sol % prod, prod


def find_contest_timestamp(schedule: List[Tuple[int, int]]) -> int:
    # check that the bus IDs are relatively primes
    if not are_numbers_relative_primes([bus_id for _, bus_id in schedule]):
        raise ValueError("Bus IDs are not relatively primes, cannot find solution")

    sol, prod = solution_two_moduli(
        -schedule[0][0], schedule[0][1], -schedule[1][0], schedule[1][1]
    )

    for i in range(2, len(schedule)):
        offset, bus_id = schedule[i]
        sol, prod = solution_two_moduli(sol, prod, -offset, bus_id)

    return sol


def main(data: str):
    timestamp, schedule = parse_input(data)

    print(
        "ID of earliest bus multiplied by waiting time:",
        find_waiting_time_multiplied_by_bus_id(timestamp, schedule),
    )
    print("Contest timestamp:", find_contest_timestamp(schedule))
