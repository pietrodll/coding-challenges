from .day12 import ShipSimulator, WaypointShipSimulator

instructions = [("F", 10), ("N", 3), ("F", 7), ("R", 90), ("F", 11)]


def assert_ship_simulator_state(sim: ShipSimulator, x: int, y: int, direction: str):
    assert sim.x == x
    assert sim.y == y
    assert sim.direction == direction


def test_ship_simulator():
    sim = ShipSimulator()
    assert_ship_simulator_state(sim, 0, 0, "E")

    expected_states = [
        (10, 0, "E"),
        (10, 3, "E"),
        (17, 3, "E"),
        (17, 3, "S"),
        (17, -8, "S"),
    ]

    for instr, state in zip(instructions, expected_states):
        sim.execute(instr)
        assert_ship_simulator_state(sim, *state)


def assert_waypoint_ship_simulator_state(
    sim: WaypointShipSimulator, x: int, y: int, dx: int, dy: int
):
    assert sim.x == x
    assert sim.y == y
    assert sim.dx == dx
    assert sim.dy == dy


def test_waypoint_ship_simulator():
    sim = WaypointShipSimulator()
    assert_waypoint_ship_simulator_state(sim, 0, 0, 10, 1)

    expected_states = [
        (100, 10, 10, 1),
        (100, 10, 10, 4),
        (170, 38, 10, 4),
        (170, 38, 4, -10),
        (214, -72, 4, -10),
    ]

    for instr, state in zip(instructions, expected_states):
        sim.execute(instr)
        assert_waypoint_ship_simulator_state(sim, *state)
