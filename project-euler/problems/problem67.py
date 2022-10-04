"""
Problem 67
==========
"""

from utils.triangle import max_path_to_bottom


def load_from_file(filename):
    file = open(filename, "r")
    L = []
    for line in file:
        L += [int(num) for num in line.split(" ")]
    return L


def main():
    L = load_from_file("data/p067_triangle.txt")
    print(max_path_to_bottom(L))
