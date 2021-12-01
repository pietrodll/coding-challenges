"""Module containing an implementation of a priority queue"""

def _parent(i):
    return (i - 1) // 2

def _right_child(i):
    return 2*i + 2

def _left_child(i):
    return 2*i + 1


class PriorityQueue():
    """This class represents a priority queue"""

    def __init__(self):
        self.items = []


    def __len__(self):
        return len(self.items)


    def __iter__(self):
        return iter(self.items)

    def __contains__(self, key):
        return key in self.items


    def empty(self):
        return self.items == []


    def insert(self, k):
        self.items.append(k)
        self._percolate_up()


    def pop(self):
        self._swap(0, len(self.items)-1)
        value = self.items.pop()
        self._percolate_down()
        return value

    def get_min(self):
        return self.items[0]


    def _percolate_up(self):
        i = len(self.items) - 1
        while _parent(i) >= 0 and self.items[i] < self.items[_parent(i)]:
            self._swap(i, _parent(i))


    def _swap(self, i, j):
        self.items[i], self.items[j] = self.items[j], self.items[i]


    def _min_child(self, i):
        if _right_child(i) >= len(self.items):
            return _left_child(i)
        if self.items[_right_child(i)] < self.items[_left_child(i)]:
            return _right_child(i)
        return _left_child(i)


    def _percolate_down(self):
        i = 0
        while self._min_child(i) < len(self.items):
            mc = self._min_child(i)
            self._swap(i, mc)
            i = mc


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
