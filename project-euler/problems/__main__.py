"""Main module to run all problems"""

import sys
import argparse
import importlib


def parse_problem_from_args():
    parser = argparse.ArgumentParser(
        prog="problems", description="Parse problem to run"
    )

    parser.add_argument(
        "problem", metavar="p", type=int, help="The number of the problem to run",
    )

    args = parser.parse_args()
    return args.problem


def run_problem(problem: int):
    print(f"--------- PROBLEM {problem} ---------")

    try:
        module = importlib.import_module(f"problems.problem{problem}")
        module.main()

    except ImportError:
        print(f"Solution for problem {problem} not yet implements")
        sys.exit(1)


def main():
    problem = parse_problem_from_args()
    run_problem(problem)


if __name__ == "__main__":
    main()
