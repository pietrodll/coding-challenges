"""Day 11"""

from typing import Dict, List, Tuple, Generator


def parse_input(data: str) -> List[List[int]]:
    mapping = {".": 0, "L": 1, "#": 2}

    def map_func(c):
        return mapping[c]

    return [list(map(map_func, line)) for line in data.splitlines()]


def get_adjacent_seats(
    i: int, j: int, grid: List[List[int]]
) -> Generator[Tuple[int, int], None, None]:
    height = len(grid)
    width = len(grid[0])

    if i > 0:
        yield (i - 1, j)

        if j > 0:
            yield (i - 1, j - 1)

        if j < width - 1:
            yield (i - 1, j + 1)

    if i < height - 1:
        yield (i + 1, j)

        if j > 0:
            yield (i + 1, j - 1)

        if j < width - 1:
            yield (i + 1, j + 1)

    if j > 0:
        yield (i, j - 1)

    if j < width - 1:
        yield (i, j + 1)


def get_visible_seats(
    i: int, j: int, grid: List[List[int]]
) -> Generator[Tuple[int, int], None, None]:
    height = len(grid)
    width = len(grid[0])
    directions = []

    if i > 0:
        # go up
        directions.append((-1, 0))

        if j > 0:
            # go up left
            directions.append((-1, -1))

        if j < width - 1:
            # go up right
            directions.append((-1, 1))

    if i < height - 1:
        # go down
        directions.append((1, 0))

        if j > 0:
            # go down left
            directions.append((1, -1))

        if j < width - 1:
            # go down right
            directions.append((1, 1))

    if j > 0:
        # go left
        directions.append((0, -1))

    if j < width - 1:
        # go right
        directions.append((0, 1))

    for pi, pj in directions:
        x = i + pi
        y = j + pj

        while (
            (pi == 0 or 0 < x < height - 1)
            and (pj == 0 or 0 < y < width - 1)
            and grid[x][y] == 0
        ):
            x += pi
            y += pj

        if grid[x][y] != 0:
            yield (x, y)


class SeatSimulator:
    def __init__(self, grid: List[List[int]]):
        self.grid = grid
        self.height = len(grid)
        self.width = len(grid[0])
        self.close_seats = self._init_close_seats()

    def _init_close_seats(self) -> Dict[Tuple[int, int], List[Tuple[int, int]]]:
        """Defines the mapping from a seat to its close seats (i.e. the seats to check
        in order to know if the seat becomes free/occupied)
        """

        close_seats = {}

        for i in range(self.height):
            for j in range(self.width):
                close_seats[(i, j)] = list(
                    filter(
                        lambda pos: self._get(*pos) != 0,
                        get_adjacent_seats(i, j, self.grid),
                    )
                )

        return close_seats

    def _get(self, i: int, j: int):
        return self.grid[i][j]

    def _becomes_occupied(self, i: int, j: int) -> bool:
        """An empty seat becomes occupied if there are no occupied seats around"""

        if self.grid[i][j] != 1:
            return False

        close_seats = self.close_seats[(i, j)]
        return not any(self._get(a, b) == 2 for a, b in close_seats)

    def _becomes_empty(self, i: int, j: int) -> bool:
        """An occupied seat becomes empty if there are 4 or more occupied seats
        around"""

        if self.grid[i][j] != 2:
            return False

        close_seats = self.close_seats[(i, j)]
        return sum(self.grid[a][b] == 2 for a, b in close_seats) >= 4

    def next_state(self):
        next_grid = [[0] * self.width for _ in range(self.height)]
        has_changed = False

        for i in range(self.height):
            for j in range(self.width):
                if self._becomes_occupied(i, j):
                    next_grid[i][j] = 2
                    has_changed = True

                elif self._becomes_empty(i, j):
                    next_grid[i][j] = 1
                    has_changed = True

                else:
                    next_grid[i][j] = self._get(i, j)

        self.grid = next_grid
        return has_changed

    def compute_until_equilibrium(self) -> int:
        steps = 0

        while self.next_state():
            steps += 1

        return steps

    def count_occupied_seats(self) -> int:
        occupied_cnt = 0

        for i in range(self.height):
            for j in range(self.width):
                if self._get(i, j) == 2:
                    occupied_cnt += 1

        return occupied_cnt


class SeatSimulatorV2(SeatSimulator):
    def _init_close_seats(self) -> Dict[Tuple[int, int], List[Tuple[int, int]]]:
        close_seats = {}

        for i in range(self.height):
            for j in range(self.width):
                close_seats[(i, j)] = list(get_visible_seats(i, j, self.grid))

        return close_seats

    def _becomes_empty(self, i: int, j: int) -> bool:
        """An occupied seat becomes empty if there are 5 or more occupied seats
        around"""

        if self.grid[i][j] != 2:
            return False

        close_seats = self.close_seats[(i, j)]
        return sum(self.grid[a][b] == 2 for a, b in close_seats) >= 5


def main(data: str):
    seats = parse_input(data)

    sim = SeatSimulator(seats)
    sim.compute_until_equilibrium()
    print("Number of occupied seats:", sim.count_occupied_seats())

    sim2 = SeatSimulatorV2(seats)
    sim2.compute_until_equilibrium()
    print("Number of occupied seats (part 2):", sim2.count_occupied_seats())
