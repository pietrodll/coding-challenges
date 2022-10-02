from .day8 import (
    execute_program_until_infinite_loop,
    find_faulty_instruction,
    parse_input,
)

data = """nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6"""


def test_parse_input():
    expected = [
        ("nop", 0),
        ("acc", 1),
        ("jmp", 4),
        ("acc", 3),
        ("jmp", -3),
        ("acc", -99),
        ("acc", 1),
        ("jmp", -4),
        ("acc", 6),
    ]

    assert parse_input(data) == expected


def test_execute_program_until_infinite_loop():
    instructions = parse_input(data)

    assert execute_program_until_infinite_loop(instructions) == 5


def test_find_faulty_instruction():
    instructions = parse_input(data)

    assert find_faulty_instruction(instructions) == 8
