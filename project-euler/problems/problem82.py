"""
Problem 82
==========
"""

from utils.priority_queue import PriorityQueue


def load_matrix(filename):
    with open(filename, "r") as file:
        matrix = []

        for line in file:
            int_line = [int(x) for x in line.split(",")]
            matrix.append(int_line)

        return matrix


def get_children(i, j, n):
    if i == 0:
        if j == n - 1:
            return [(i + 1, j)]

        return [(i + 1, j), (i, j + 1)]

    if i == n - 1:
        if j == n - 1:
            return [(i - 1, j)]

        return [(i, j + 1), (i - 1, j)]

    if j == n - 1:
        return [(i - 1, j), (i + 1, j)]

    return [(i, j + 1), (i - 1, j), (i + 1, j)]


def tab_index(i, j, n):
    return i * n + j


def matrix_to_graph(M):
    n = len(M)
    G = []

    for i in range(n):
        for j in range(n):
            children = []

            for (a, b) in get_children(i, j, n):
                children.append((tab_index(a, b, n), M[a][b]))

            G.append(children)

    return G


def distance_to_node(G, s):
    dist = [float("inf")] * len(G)
    dist[s] = 0

    visited = set()

    queue = PriorityQueue(prio_func=lambda x: dist[x])
    queue.insert(s)

    while not queue.empty():
        u = queue.pop()

        for v, w in G[u]:
            if v not in queue and v not in visited:
                queue.insert(v)

            if dist[v] > dist[u] + w:
                dist[v] = dist[u] + w

                if v in queue:
                    queue.decrease_priority(v)

        visited.add(u)

    return dist


def distance_to_last_col(G, s, n):
    D = distance_to_node(G, s)
    return min([D[tab_index(i, n - 1, n)] for i in range(n)])


def min_distance_to_col(M):
    n = len(M)
    G = matrix_to_graph(M)
    paths = [0] * n
    for i in range(n):
        d = distance_to_last_col(G, tab_index(i, 0, n), n)
        print(d)
        paths[i] = M[i][0] + d
    return min(paths)


def main():
    matrix = load_matrix("data/p082_matrix.txt")

    # print(matrix)

    # G = matrix_to_graph(matrix)
    # print(G)
    # print(distance_to_last_col(G, 80, 80))

    print(min_distance_to_col(matrix))
