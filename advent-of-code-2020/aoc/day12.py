"""Day 12"""


from typing import List, Tuple


def parse_input(data: str) -> List[Tuple[str, int]]:
    instructions = []

    for instruction_str in data.splitlines():
        instructions.append((instruction_str[0], int(instruction_str[1:])))

    return instructions


class ShipSimulator:
    def __init__(self) -> None:
        self.x = 0
        self.y = 0
        self.direction = "E"

    def move(self, direction: str, distance: int):
        if direction == "N":
            self.y += distance
        elif direction == "S":
            self.y -= distance
        elif direction == "W":
            self.x -= distance
        else:
            self.x += distance

    def forward(self, distance: int):
        self.move(self.direction, distance)

    def turn(self, direction: str, degrees: int):
        if direction == "R":
            sequence = ["N", "E", "S", "W"]
        else:
            sequence = ["N", "W", "S", "E"]

        index = sequence.index(self.direction)
        next_index = (index + degrees // 90) % 4
        self.direction = sequence[next_index]

    def execute(self, instruction: Tuple[str, int]):
        inst, value = instruction

        if inst == "F":
            self.forward(value)
        elif inst == "L" or inst == "R":
            self.turn(inst, value)
        else:
            self.move(inst, value)

    def execute_multi(self, instructions: List[Tuple[str, int]]):
        for instruction in instructions:
            self.execute(instruction)


class WaypointShipSimulator(ShipSimulator):
    def __init__(self) -> None:
        super().__init__()
        self.dx = 10
        self.dy = 1

    def move(self, direction: str, distance: int):
        if direction == "N":
            self.dy += distance
        elif direction == "S":
            self.dy -= distance
        elif direction == "W":
            self.dx -= distance
        else:
            self.dx += distance

    def turn(self, direction: str, degrees: int):
        cos = [1, 0, -1, 0]
        sin = [0, 1, 0, -1]

        cosVal = cos[(degrees // 90) % 4]
        sinVal = sin[(degrees // 90) % 4]

        if direction == "R":
            sinVal = -sinVal

        self.dx, self.dy = (
            cosVal * self.dx - sinVal * self.dy,
            sinVal * self.dx + cosVal * self.dy,
        )

    def forward(self, distance: int):
        self.x += self.dx * distance
        self.y += self.dy * distance


def manhattan_distance(instructions: List[Tuple[str, int]], waypoint=False) -> int:
    sim = WaypointShipSimulator() if waypoint else ShipSimulator()

    sim.execute_multi(instructions)

    return abs(sim.x) + abs(sim.y)


def main(data: str):
    instructions = parse_input(data)

    print("Manhattan distance:", manhattan_distance(instructions))

    print(
        "Manhattan distance (with waypoint):",
        manhattan_distance(instructions, waypoint=True),
    )
