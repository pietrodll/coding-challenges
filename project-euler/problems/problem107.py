"""
Problem 107
===========
"""

from utils.priority_queue import Node, PriorityQueue


def load_matrix(filename):
    M = []

    with open(filename, "r") as f:
        for line in f:
            L = []

            for c in line.split(","):
                if c not in ("-", "-\n"):
                    L.append(int(c))

                else:
                    L.append(0)

            M.append(L)

        return M


def matrix_to_graph(M):
    G = []

    for line in M:
        neighbors = []

        for i, x in enumerate(line):
            if x != 0:
                neighbors.append((i, x))
        G.append(neighbors)
    return G


def prim(G, s):
    n = len(G)
    d = [float("inf")] * n
    d[s] = 0
    p = list(range(n))
    H = {}
    F = PriorityQueue()
    F.insert(Node(s, d[s]))
    E = []

    for _ in range(n):
        U = F.pop()

        if U.value != s:
            E.append((U.value, p[U.value]))

        for v, w in G[U.value]:
            V = Node(v, w)

            if V not in F and v not in H:
                if d[v] > w:
                    d[v] = w
                    V.priority = w
                    p[v] = U.value

                F.insert(V)

            elif d[v] > w:
                d[v] = w
                V.priority = w
                p[v] = U.value

        H[U.value] = True

    return E


def total_weight(G):
    W = 0

    for x in G:
        for (_, w) in x:
            W += w

    return W // 2


def max_saving(M):
    G = matrix_to_graph(M)
    W = total_weight(G)
    E = prim(G, 0)
    w = 0

    for (i, j) in E:
        w += M[i][j]

    return W - w


def main():
    M = [
        [0, 16, 12, 21, 0, 0, 0],
        [16, 0, 0, 17, 20, 0, 0],
        [12, 0, 0, 28, 0, 31, 0],
        [21, 17, 28, 0, 18, 19, 23],
        [0, 20, 0, 18, 0, 0, 11],
        [0, 0, 31, 19, 0, 0, 27],
        [0, 0, 0, 23, 11, 27, 0],
    ]

    # M = load_matrix('data/p107_network.txt')

    G = matrix_to_graph(M)
    E = prim(G, 0)
    print(G)
    print(E)
    print(total_weight(G))
    print(max_saving(M))
