from .day7 import (
    count_containing_possibilities,
    count_total_inside,
    parse_contained_relations,
    parse_containing_relations,
    parse_input,
)

data = """light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags."""


def test_parse_input():
    expected = [
        ("light red", [(1, "bright white"), (2, "muted yellow")]),
        ("dark orange", [(3, "bright white"), (4, "muted yellow")]),
        ("bright white", [(1, "shiny gold")]),
        ("muted yellow", [(2, "shiny gold"), (9, "faded blue")]),
        ("shiny gold", [(1, "dark olive"), (2, "vibrant plum")]),
        ("dark olive", [(3, "faded blue"), (4, "dotted black")]),
        ("vibrant plum", [(5, "faded blue"), (6, "dotted black")]),
        ("faded blue", []),
        ("dotted black", []),
    ]

    assert parse_input(data) == expected


def test_parse_contained_relations():
    expected = {
        "light red": [],
        "dark orange": [],
        "bright white": [(1, "light red"), (3, "dark orange")],
        "muted yellow": [(2, "light red"), (4, "dark orange")],
        "shiny gold": [(1, "bright white"), (2, "muted yellow")],
        "dark olive": [(1, "shiny gold")],
        "vibrant plum": [(2, "shiny gold")],
        "faded blue": [(9, "muted yellow"), (3, "dark olive"), (5, "vibrant plum")],
        "dotted black": [(4, "dark olive"), (6, "vibrant plum")],
    }

    assert parse_contained_relations(parse_input(data)) == expected


def test_count_containing_possibilities():
    graph = parse_contained_relations(parse_input(data))

    assert count_containing_possibilities(graph, "shiny gold") == 4


def test_parse_containing_relations():
    expected = {
        "light red": [(1, "bright white"), (2, "muted yellow")],
        "dark orange": [(3, "bright white"), (4, "muted yellow")],
        "bright white": [(1, "shiny gold")],
        "muted yellow": [(2, "shiny gold"), (9, "faded blue")],
        "shiny gold": [(1, "dark olive"), (2, "vibrant plum")],
        "dark olive": [(3, "faded blue"), (4, "dotted black")],
        "vibrant plum": [(5, "faded blue"), (6, "dotted black")],
        "faded blue": [],
        "dotted black": [],
    }

    assert parse_containing_relations(parse_input(data)) == expected


def test_count_total_inside():
    containing = parse_containing_relations(parse_input(data))

    assert count_total_inside(containing, "shiny gold") == 32
