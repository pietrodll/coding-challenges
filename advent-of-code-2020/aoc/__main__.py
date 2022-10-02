"""Main module to run the scripts for all days"""

from .utils import run_day, parse_days_from_args


def main():
    for day in parse_days_from_args():
        run_day(day)


if __name__ == "__main__":
    main()
