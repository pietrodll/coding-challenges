"""
Problem 107
===========
"""

from typing import List, Tuple
from utils.priority_queue import PriorityQueue


def load_matrix(filename: str) -> List[List[int]]:
    """Loads a matrix from a text file with the format used in the problem"""

    with open(filename, "r") as f:
        M = []

        for line in f:
            L = []

            for c in line.split(","):
                if c not in ("-", "-\n"):
                    L.append(int(c))

                else:
                    L.append(None)

            M.append(L)

        return M


def matrix_to_graph(M: List[List[int]]) -> List[List[Tuple[int, int]]]:
    graph = []

    for line in M:
        neighbors = []

        for i, x in enumerate(line):
            if x is not None:
                neighbors.append((i, x))

        graph.append(neighbors)

    return graph


def min_spanning_tree(
    graph: List[List[Tuple[int, int]]], start: int
) -> List[Tuple[int, int]]:
    n = len(graph)

    cost = [float("inf")] * n
    cost[start] = 0

    parent = list(range(n))
    finished = set()
    queue = PriorityQueue(prio_func=lambda x: cost[x])
    queue.insert(start)
    edges = []

    while not queue.empty():
        u = queue.pop()

        if u != start:
            edges.append((u, parent[u]))

        for v, w in graph[u]:
            if v not in queue and v not in finished:
                queue.insert(v)

            if cost[v] > w:
                cost[v] = w
                parent[v] = u

                if v in queue:
                    queue.decrease_priority(v)

        finished.add(u)

    return edges


def graph_weight(graph: List[List[Tuple[int, int]]]) -> int:
    """Computes the total weight of a graph"""
    W = 0

    for x in graph:
        for _, w in x:
            W += w

    return W // 2


def edges_weight(matrix: List[List[int]], edges: List[Tuple[int, int]]) -> int:
    return sum(matrix[i][j] for i, j in edges)


def max_saving(matrix: List[List[int]]):
    """Computes the maximum possible saving achievable by removing redundant edges"""

    G = matrix_to_graph(matrix)
    max_weight = graph_weight(G)
    E = min_spanning_tree(G, 0)
    min_weight = edges_weight(matrix, E)

    return max_weight - min_weight


def main():
    # M = [
    #     [None, 16, 12, 21, None, None, None],
    #     [16, None, None, 17, 20, None, None],
    #     [12, None, None, 28, None, 31, None],
    #     [21, 17, 28, None, 18, 19, 23],
    #     [None, 20, None, 18, None, None, 11],
    #     [None, None, 31, 19, None, None, 27],
    #     [None, None, None, 23, 11, 27, None],
    # ]

    M = load_matrix("data/p107_network.txt")

    print(max_saving(M))
