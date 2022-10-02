"""Day 8"""

import re
from typing import List, Tuple


def parse_input(data: str) -> List[Tuple[str, int]]:
    instruction_pattern = re.compile(r"(nop|acc|jmp) ((?:\+|-)\d+)")
    instructions = []

    for instruction_text in data.splitlines():
        match = instruction_pattern.fullmatch(instruction_text)

        if not match:
            raise ValueError(f"Cannot parse {instruction_text}")

        instructions.append((match.group(1), int(match.group(2))))

    return instructions


def execute_program(instructions: List[Tuple[str, int]]):
    executed_instructions = set()
    i = 0
    acc = 0

    while i < len(instructions):
        if i in executed_instructions:
            raise TimeoutError(
                f"Infinite loop! Current state: acc = {acc}; i = {i}", acc, i
            )

        executed_instructions.add(i)
        instruction_type, value = instructions[i]

        if instruction_type == "acc":
            acc += value
            i += 1

        elif instruction_type == "jmp":
            i += value

        else:
            i += 1

    return acc


def execute_program_until_infinite_loop(instructions: List[Tuple[str, int]]):
    try:
        execute_program(instructions)

    except TimeoutError as e:
        return int(e.args[1])


def find_faulty_instruction(instructions: List[Tuple[str, int]]):
    for i, (instruction, value) in enumerate(instructions):
        if instruction == "jmp" or (instruction == "nop" and value != 0):
            copy = list(instructions)

            if instruction == "jmp":
                copy[i] = ("nop", value)

            else:
                copy[i] = ("jmp", value)

            try:
                return execute_program(copy)

            except TimeoutError:
                pass

    raise ValueError("Cannot find faulty instruction")


def main(data: str):
    instructions = parse_input(data)

    print(
        "Executing program until infinite loop. ACC =",
        execute_program_until_infinite_loop(instructions),
    )

    print(
        "Result after fixing faulty instruction:", find_faulty_instruction(instructions)
    )
