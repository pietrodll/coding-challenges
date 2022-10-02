"""Day 16"""

import re
from typing import Dict, List, Tuple


RuleType = Tuple[str, List[Tuple[int, int]]]
TicketType = List[int]


def parse_rules(rules_block: str) -> List[RuleType]:
    rule_pattern = re.compile(r"([^:]+):(.+)")
    range_pattern = re.compile(r"(\d+)-(\d+)")

    rules = []
    for rule_str in rules_block.splitlines():
        match = rule_pattern.fullmatch(rule_str)

        if not match:
            raise ValueError(f'Cannot parse rule "{rule_str}"')

        field = match.group(1)
        ranges = []

        for range_str in match.group(2).strip().split(" or "):
            range_match = range_pattern.fullmatch(range_str)

            if not range_match:
                raise ValueError(f'Cannot parse rule "{rule_str}"')

            ranges.append((int(range_match.group(1)), int(range_match.group(2))))

        rules.append((field, ranges))

    return rules


def parse_ticket(ticket_block: str) -> List[int]:
    ticket_lines = ticket_block.splitlines()

    if ticket_lines[0] != "your ticket:":
        raise ValueError("Invalid ticket input")

    return list(map(int, ticket_lines[1].split(",")))


def parse_nearby_tickets(nearby_tickets_block: str) -> List[List[int]]:
    lines = nearby_tickets_block.splitlines()

    if lines[0] != "nearby tickets:":
        raise ValueError("Invalid nearby tickets input")

    nearby_tickets = []

    for line in lines[1:]:
        nearby_tickets.append(list(map(int, line.split(","))))

    return nearby_tickets


def parse_input(data: str) -> Tuple[List[RuleType], TicketType, List[TicketType]]:
    blocks = data.split("\n\n")

    rules = parse_rules(blocks[0])
    ticket = parse_ticket(blocks[1])
    nearby_tickets = parse_nearby_tickets(blocks[2])

    return rules, ticket, nearby_tickets


def is_valid_value_for_rule(value: int, rule: RuleType) -> bool:
    return any(value >= inf and value <= sup for inf, sup in rule[1])


def is_valid_value(value: int, rules: List[RuleType]) -> bool:
    return any(is_valid_value_for_rule(value, rule) for rule in rules)


def compute_ticket_scanning_error_rate(
    tickets: List[TicketType], rules: List[RuleType]
) -> int:
    rate = 0

    for ticket in tickets:
        for value in ticket:
            if not is_valid_value(value, rules):
                rate += value

    return rate


def is_valid_ticket(ticket: TicketType, rules: List[RuleType]) -> bool:
    return all(is_valid_value(value, rules) for value in ticket)


def decode_ticket(
    rules: List[RuleType], ticket: TicketType, nearby_tickets: List[TicketType]
) -> Dict[str, int]:
    valid_tickets = [t for t in nearby_tickets if is_valid_ticket(t, rules)]

    all_fields = {rule[0] for rule in rules}
    possible_fields_by_position = [all_fields.copy() for _ in range(len(ticket))]

    for valid_ticket in valid_tickets:
        for pos, value in enumerate(valid_ticket):
            possible_fields = {
                rule[0] for rule in rules if is_valid_value_for_rule(value, rule)
            }
            possible_fields_by_position[pos].intersection_update(possible_fields)

    decrypted_ticket = {}

    possible_fields_with_position = list(enumerate(possible_fields_by_position))
    possible_fields_with_position.sort(key=lambda tup: len(tup[1]))

    # there should be at least one position that can be mapped to only one field
    # we start by this one, then use it to

    for i, (pos, possible_fields) in enumerate(possible_fields_with_position):
        if len(possible_fields) > 1:
            raise ValueError(
                "Could not decrypt ticket, "
                "there are multiple possible fields for position {}: {}".format(
                    pos, ", ".join(possible_fields)
                )
            )

        for j in range(i + 1, len(possible_fields_with_position)):
            possible_fields_with_position[j][1].difference_update(possible_fields)

        decrypted_ticket[possible_fields.pop()] = ticket[pos]

    return decrypted_ticket


def decript_and_multiply_departure_values(
    rules: List[RuleType], ticket: TicketType, nearby_tickets: List[TicketType]
) -> int:
    decrypted_ticket = decode_ticket(rules, ticket, nearby_tickets)

    prod = 1

    for field in decrypted_ticket:
        if field.startswith("departure"):
            prod *= decrypted_ticket[field]

    return prod


def main(data: str):
    rules, ticket, nearby_tickets = parse_input(data)

    print(
        "Ticket scanning error rate:",
        compute_ticket_scanning_error_rate(nearby_tickets, rules),
    )
    print(
        "Departure values multiplied together:",
        decript_and_multiply_departure_values(rules, ticket, nearby_tickets),
    )
