"""Utility functions"""

import argparse
from contextlib import contextmanager
from importlib import import_module
from pathlib import Path


@contextmanager
def open_input(day: int):
    filepath = Path(__file__).parent.parent / f"input/day{day}.txt"

    with open(filepath) as input_file:
        yield input_file.read().strip()


def run_day(day: int):
    print(f"--------------- DAY {day} ---------------")
    try:
        module = import_module(f"aoc.day{day}")

        with open_input(day) as data:
            module.main(data)

    except ImportError:
        print(f"Script for day {day} not yet implemented")

    print()


def parse_days_from_args():
    parser = argparse.ArgumentParser(description="Parse days to run")

    parser.add_argument("-a", "--all", action="store_true", help="Run all days")
    parser.add_argument("days", metavar="D", type=int, nargs="*")

    args = parser.parse_args()

    if args.all:
        return range(1, 26)

    return args.days
