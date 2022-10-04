"""A module containing functions to compute permutations"""


def comb(n, k):
    """A recursive function computing all the parts of [0 ... n-1] of length k"""

    if k == 1:
        return [[i] for i in range(n)]
    C = []
    for i in range(k, n):
        C += [([i] + c) for c in comb(i - 1, k - 1)]
    return C


def all_comb(N, verbose=0):
    """A function computing the permutations of [0 ... n-1] of length k for k going from 1 to n"""

    L = [[[]] * N for i in range(N)]

    for n in range(N):
        L[n][0] = [[i] for i in range(n + 1)]

    for k in range(1, N):
        for n in range(k, N):
            L[n][k] = [([n] + C) for C in L[n - 1][k - 1]] + L[n - 1][k]

    if verbose == 1:
        print(print_comb(L))

    return L[-1]


def print_comb(L):
    """Display the permutation matrix in a nice way"""

    s = ""
    for line in L:
        sline = ""
        for part_list in line:
            spart = "[ "
            for part in part_list:
                part = [str(i) for i in part]
                spart += "{" + ", ".join(part) + "}, "
            sline += spart + "]\t"
        s += sline + "\n"
    return s


def permutations(N):
    """Computes all the permutations of [0 ... n-1]"""

    def aux(elements):
        if len(elements) == 1:
            return [[elements.pop()]]

        perms = []

        for x in elements:
            for perm in aux(elements - {x}):
                perms.append([x] + perm)

        return perms

    return aux(set(range(N)))


def iter_permutations(L):
    """Iterates through the permutations of L. The permutations are generated after removing the
    duplicates of L. The elements of L must be hashable"""

    def aux(elements):
        if len(elements) == 1:
            yield [elements.pop()]

        for x in elements:
            for perm in aux(elements - {x}):
                yield [x] + perm

    return aux(set(L))
