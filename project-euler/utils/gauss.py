"""Module containing functions to implement the gauss pivot algorithm"""


def transvection(A, i, j, mu):
    """We add the j-th line of A multiplied by mu to the i-th line of A"""

    for k in range(len(A)):
        A[i][k] += mu * A[j][k]


def swap_lines(A, i, j):
    """Swaps two lines of the matrix"""

    A[i], A[j] = A[j], A[i]


def pivot_index(A, i0):
    """Finds a line of index >= i with the highest coefficient at column i in absolute value"""

    i_max = i0

    for i in range(i0 + 1, len(A)):
        if abs(A[i]) > abs(A[i_max]):
            i_max = i

    return i_max


def gauss(A0, Y0):
    """Solves the system A0*X = Y0"""

    # copy the matrix and the vector
    n = len(A0)
    A = [A0[i][:] for i in range(n)]
    Y = [[Y0[i]] for i in range(n)]

    for i in range(n):

        # find the pivot and swap lines if necessary
        j = pivot_index(A, i)
        if j != i:
            swap_lines(A, i, j)
            swap_lines(Y, i, j)

        # transvections
        for j in range(i+1, n):
            mu = A[i][j] / A[i][i]
            transvection(A, i, j, mu)
            transvection(Y, i, j, mu)

    # compute the solution
    X = [0] * n
    for i in range(n-1, -1, -1):
        X[i] = 1 / A[i][i] * (Y[i][0] - sum([A[i][j] * X[j] for j in range(j+1, n)]))

    return X
