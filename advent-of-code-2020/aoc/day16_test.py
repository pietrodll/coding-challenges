from .day16 import (
    compute_ticket_scanning_error_rate,
    decode_ticket,
    is_valid_value,
    parse_input,
)

data1 = """class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12"""

data2 = """class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9"""


def test_parse_input():
    rules, ticket, nearby_tickets = parse_input(data2)

    assert rules == [
        ("class", [(0, 1), (4, 19)]),
        ("row", [(0, 5), (8, 19)]),
        ("seat", [(0, 13), (16, 19)]),
    ]
    assert ticket == [11, 12, 13]
    assert nearby_tickets == [[3, 9, 18], [15, 1, 5], [5, 14, 9]]


def test_is_valid_value():
    rules, *_ = parse_input(data1)

    assert is_valid_value(40, rules) is True
    assert is_valid_value(4, rules) is False
    assert is_valid_value(55, rules) is False
    assert is_valid_value(12, rules) is False


def test_compute_ticket_scanning_error_rate():
    rules, _, nearby_tickets = parse_input(data1)
    assert compute_ticket_scanning_error_rate(nearby_tickets, rules) == 71


def test_decode_ticket():
    rules, ticket, nearby_tickets = parse_input(data2)
    assert decode_ticket(rules, ticket, nearby_tickets) == {
        "class": 12,
        "row": 11,
        "seat": 13,
    }
