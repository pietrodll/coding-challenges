"""
Problem 14
==========
"""


def collatz_sequence(n):
    u = n
    c = 1
    while u != 1:
        c += 1
        if u % 2 == 0:
            u = u // 2
        else:
            u = 3 * u + 1
    return c


def collatz_list_naive(N):
    L = [0, 1]
    for i in range(2, N):
        L.append(collatz_sequence(i))
    return L


def collatz_list(N):
    L = [0, 1]
    for i in range(2, N):
        u = i
        c = 1
        while u != 1:
            c += 1
            if u % 2 == 0:
                u = u // 2
                if u < len(L):
                    c += L[u] - 1
                    u = 1
            else:
                u = 3 * u + 1
        L.append(c)
    return L


def longest_collatz(N):
    L = [0, 1]
    m = 1
    i_max = 1
    for i in range(2, N):
        u = i
        c = 1
        while u != 1:
            c += 1
            if u % 2 == 0:
                u = u // 2
                if u < len(L):
                    c += L[u] - 1
                    u = 1
            else:
                u = 3 * u + 1
        L.append(c)
        if c > m:
            m = c
            i_max = i
    return i_max, m


def main():
    l1 = collatz_list_naive(20)
    l2 = collatz_list(20)

    print(l1)
    print(l2)

    print(longest_collatz(1000000))
