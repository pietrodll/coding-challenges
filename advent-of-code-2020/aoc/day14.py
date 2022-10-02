"""Day 14"""

import re
from typing import List, Protocol


class ProgramState:
    def __init__(self) -> None:
        self.memory = {}
        self.mask = None


class Instruction(Protocol):
    def update_state(self, state: ProgramState):
        ...

    def update_state_v2(self, state: ProgramState):
        ...


class UpdateMaskInstruction:
    def __init__(self, mask: str) -> None:
        self.mask = mask

    def update_state(self, state: ProgramState):
        state.mask = self.mask

    def update_state_v2(self, state: ProgramState):
        state.mask = self.mask

    def __eq__(self, __o: object) -> bool:
        if isinstance(self, __o.__class__):
            return self.mask == __o.mask

        return False


def apply_mask(mask: str, value: int) -> int:
    result = 0
    power_of_2 = 1

    for bit in mask[::-1]:
        if bit == "X":
            result += (value % 2) * power_of_2
        else:
            result += int(bit) * power_of_2

        power_of_2 = power_of_2 << 1
        value = value >> 1

    return result


def apply_mask_v2(mask: str, value: int) -> List[int]:
    results = [0]
    power_of_2 = 1

    for bit in mask[::-1]:
        if bit == "X":
            results += [result + power_of_2 for result in results]
        elif bit == "0":
            for i in range(len(results)):
                results[i] += (value % 2) * power_of_2
        else:
            for i in range(len(results)):
                results[i] += power_of_2

        power_of_2 = power_of_2 << 1
        value = value >> 1

    return results


class AssignValueInstruction:
    def __init__(self, address: int, decimal_value: int) -> None:
        self.address = address
        self.decimal_value = decimal_value

    def update_state(self, state: ProgramState):
        state.memory[self.address] = apply_mask(state.mask, self.decimal_value)

    def update_state_v2(self, state: ProgramState):
        for address in apply_mask_v2(state.mask, self.address):
            state.memory[address] = self.decimal_value

    def __eq__(self, __o: object) -> bool:
        if isinstance(self, __o.__class__):
            return (
                self.address == __o.address and self.decimal_value == __o.decimal_value
            )

        return False


def parse_input(data: str) -> List[Instruction]:
    mask_pattern = re.compile(r"mask = ((?:0|1|X)+)")
    assignment_pattern = re.compile(r"mem\[(\d+)\] = (\d+)")

    instructions = []

    for line in data.splitlines():
        match = mask_pattern.fullmatch(line)

        if match:
            instructions.append(UpdateMaskInstruction(match.group(1)))

        else:
            match = assignment_pattern.fullmatch(line)

            if match:
                instructions.append(
                    AssignValueInstruction(int(match.group(1)), int(match.group(2)))
                )

            else:
                raise ValueError(f'Cannot parse line "{line}"')

    return instructions


def execute_instructions(instructions: List[Instruction], v2=False) -> ProgramState:
    state = ProgramState()

    for instruction in instructions:
        if v2:
            instruction.update_state_v2(state)
        else:
            instruction.update_state(state)

    return state


def sum_all_recorded_values(state: ProgramState):
    return sum(state.memory.values())


def main(data: str):
    instructions = parse_input(data)

    print(
        "Sum of values in memory:",
        sum_all_recorded_values(execute_instructions(instructions)),
    )
    print(
        "Sum of values in memory (V2):",
        sum_all_recorded_values(execute_instructions(instructions, v2=True)),
    )
