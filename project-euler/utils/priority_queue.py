"""Module containing an implementation of a priority queue"""


def _parent(i):
    return (i - 1) // 2


def _right_child(i):
    return 2 * i + 2


def _left_child(i):
    return 2 * i + 1


class PriorityQueue:
    """This class represents a priority queue"""

    def __init__(self, prio_func=None):
        self.items = []
        self.positions = {}

        if prio_func is None:
            self.prio_func = lambda x: x
        else:
            self.prio_func = prio_func

    def _get_prio_at(self, position):
        return self.prio_func(self.items[position])

    def __len__(self):
        return len(self.items)

    def __iter__(self):
        return iter(self.items)

    def __contains__(self, obj):
        return obj in self.positions

    def empty(self):
        return self.items == []

    def insert(self, k):
        self.items.append(k)
        self.positions[k] = len(self.items) - 1
        self._percolate_up(len(self.items) - 1)

    def pop(self):
        self._swap(0, len(self.items) - 1)
        value = self.items.pop()
        self.positions.pop(value)
        self._percolate_down(0)
        return value

    def get_min(self):
        return self.items[0]

    def decrease_priority(self, item):
        self._percolate_up(self.positions[item])

    def increase_priority(self, item):
        self._percolate_down(self.positions[item])

    def _percolate_up(self, start_position):
        i = start_position

        while _parent(i) >= 0 and self._get_prio_at(i) < self._get_prio_at(_parent(i)):
            self._swap(i, _parent(i))
            i = _parent(i)

    def _swap(self, i, j):
        self.items[i], self.items[j] = self.items[j], self.items[i]
        self.positions[self.items[i]] = i
        self.positions[self.items[j]] = j

    def _min_child(self, i):
        if _right_child(i) >= len(self.items):
            return _left_child(i)

        if self._get_prio_at(_right_child(i)) < self._get_prio_at(_left_child(i)):
            return _right_child(i)

        return _left_child(i)

    def _percolate_down(self, start_position):
        i, mc = start_position, self._min_child(start_position)

        while mc < len(self.items) and self._get_prio_at(mc) < self._get_prio_at(i):
            self._swap(i, mc)
            i, mc = mc, self._min_child(mc)


class Node:
    """This class is can be used to represent priority queue nodes. The `value` parameter is used
    to check equality, while the `priority` parameter is used to check order
    """

    def __init__(self, value, priority):
        self.value = value
        self.priority = priority

    def __lt__(self, other):
        return self.priority < other.priority

    def __le__(self, other):
        return self.priority <= other.priority

    def __eq__(self, other):
        return self.value == other.value

    def __ne__(self, other):
        return self.value != other.value

    def __gt__(self, other):
        return self.priority > other.priority

    def __ge__(self, other):
        return self.priority >= other.priority

    def __hash__(self) -> int:
        return hash(self.value)
