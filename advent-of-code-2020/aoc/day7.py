"""Day 7"""

import re
from typing import Dict, List, Tuple


def parse_input(data: str) -> List[Tuple[str, List[Tuple[int, str]]]]:
    rule_pattern = re.compile(
        r"(\S+\s\S+) bags contain ((?:\d+ \S+\s\S+ bags?,?\s*)+|no other bags)\."
    )
    contains_pattern = re.compile(r"(\d+) (\S+\s\S+) bags?")

    bags = []

    for rule in data.splitlines():
        match = rule_pattern.fullmatch(rule)

        if not match:
            raise ValueError(f"Cannot parse {rule}")

        bag = match.group(1)
        contains = match.group(2)
        contained = []

        if contains != "no other bags":
            for contained_bag in re.split(r",\s+", contains):
                contains_match = contains_pattern.fullmatch(contained_bag)

                if not contains_match:
                    raise ValueError(f"Cannot parse {contained_bag}")

                contained.append(
                    (int(contains_match.group(1)), contains_match.group(2))
                )
        bags.append((bag, contained))

    return bags


def parse_contained_relations(
    rules: List[Tuple[str, List[Tuple[int, str]]]]
) -> Dict[str, List[Tuple[int, str]]]:
    """Generates a graph from a list of rules.
    An edge from X to Y with weight w means that bag X can fit w times in bag Y.
    """

    g = {}

    for bag, contained_bags in rules:
        if bag not in g:
            g[bag] = []

        for num, contained_bag in contained_bags:
            if contained_bag in g:
                g[contained_bag].append((num, bag))
            else:
                g[contained_bag] = [(num, bag)]

    return g


def count_containing_possibilities(
    contained_relations: Dict[str, List[Tuple[int, str]]], start: str
) -> int:
    count = 0
    visited = {bag: False for bag in contained_relations}

    p = start
    visited[p] = True
    to_visit = [bag for _, bag in contained_relations[p]]

    while len(to_visit) > 0:
        p = to_visit.pop()

        if not visited[p]:
            count += 1
            visited[p] = True

            for _, bag in contained_relations[p]:
                to_visit.insert(0, bag)

    return count


def parse_containing_relations(
    rules: List[Tuple[str, List[Tuple[int, str]]]]
) -> Dict[str, List[Tuple[int, str]]]:
    return {bag: contains for bag, contains in rules}


def count_total_inside(
    containing_relations: Dict[str, List[Tuple[int, str]]], start: str
) -> int:
    def aux(g, bag):
        tot = 0

        for num, contained_bag in g[bag]:
            tot += num + num * aux(g, contained_bag)

        return tot

    return aux(containing_relations, start)


def main(data: str):
    rules = parse_input(data)
    contained_relations = parse_contained_relations(rules)
    containing_relations = parse_containing_relations(rules)

    print(
        "Possibilities to contain a shiny gold bag:",
        count_containing_possibilities(contained_relations, "shiny gold"),
    )

    print(
        "Total bags in the shiny gold bag:",
        count_total_inside(containing_relations, "shiny gold"),
    )
