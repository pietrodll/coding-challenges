"""This module contains utility functions for graphs"""

from collections import deque


def distance_unweighted(G, start):
    """Computes the distance from a point of un unweighted graph to all the other points"""

    Q = deque([start])
    disc = set()
    dist = [0] * len(G)

    while len(Q) > 0:
        v = Q.pop()

        for w in G[v]:
            if w not in disc:
                disc.add(w)
                dist[w] = 1 + dist[v]
                Q.appendleft(w)

    return dist


def revert_graph(G):
    """Returns a reverted version of the graph, where all the edges are in the opposite direction"""

    rev = [[] for _ in range(len(G))]

    for v, neighbors in enumerate(G):
        for w in neighbors:
            rev[w].append(v)

    return rev
