"""
Problem 81
==========
"""

import numpy as np


def load_matrix(filename):
    file = open(filename, 'r')
    matrix = []
    for line in file:
        int_line = [int(x) for x in line.split(',')]
        matrix.append(int_line)
    file.close()
    return matrix


def minimal_path(M):
    n = len(M)
    P = np.inf * np.ones((n, n), dtype=np.int)
    P[0][0] = M[0][0]
    for i in range(1, n):
        P[i][0] = M[i][0] + P[i-1][0]
    for j in range(1, n):
        P[0][j] = M[0][j] + P[0][j-1]
    for i in range(1, n):
        for j in range(1, n):
            P[i][j] = M[i][j] + min(P[i-1][j], P[i][j-1])
    return P


def main():
    matrix = load_matrix('data/p081_matrix.txt')

    print(matrix)
    print(minimal_path(matrix))


if __name__ == "__main__":
    main()
