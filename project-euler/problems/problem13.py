"""
Problem 13
==========
"""


def load_numbers(filename):
    file = open(filename, "r")
    L = []

    for num in file:
        L.append(int(num))

    return L


def main():
    nums = load_numbers("data/p013_numbers.txt")

    print(nums)
    print(sum(nums))
