from .day14 import (
    AssignValueInstruction,
    UpdateMaskInstruction,
    apply_mask,
    apply_mask_v2,
    execute_instructions,
    parse_input,
    sum_all_recorded_values,
)

program = """mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0"""


def test_parse_input():
    expected = [
        UpdateMaskInstruction("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"),
        AssignValueInstruction(8, 11),
        AssignValueInstruction(7, 101),
        AssignValueInstruction(8, 0),
    ]

    assert parse_input(program) == expected


def test_apply_mask():
    assert apply_mask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", 11) == 73
    assert apply_mask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", 101) == 101
    assert apply_mask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", 0) == 64


def test_apply_mask_v2():
    assert apply_mask_v2("000000000000000000000000000000X1001X", 42) == [26, 27, 58, 59]
    assert apply_mask_v2("00000000000000000000000000000000X0XX", 26) == [
        16,
        17,
        18,
        19,
        24,
        25,
        26,
        27,
    ]


def test_execute_instructions():
    instructions = parse_input(program)
    state = execute_instructions(instructions)

    assert state.mask == "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"
    assert state.memory == {7: 101, 8: 64}


def test_execute_instructions_v2():
    program_v2 = """mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1"""

    instructions = parse_input(program_v2)
    state = execute_instructions(instructions, v2=True)

    assert sum_all_recorded_values(state) == 208
