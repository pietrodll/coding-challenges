"""Day 6"""

from typing import List


def parse_input(data: str) -> List[List[str]]:
    blocks = data.split("\n\n")

    group_answers = []

    for block in blocks:
        group_answers.append(block.splitlines())

    return group_answers


def count_group_different_answers(group_answers: List[str]) -> int:
    different_answers = set()

    for answer in group_answers:
        different_answers.update(answer)

    return len(different_answers)


def count_group_all_answers(group_answers: List[str]) -> int:
    all_answered = set(group_answers[0])

    for answer in group_answers[1:]:
        all_answered.intersection_update(answer)

    return len(all_answered)


def main(data: str):
    groups_answers = parse_input(data)

    total_count = sum(map(count_group_different_answers, groups_answers))
    all_answered_count = sum(map(count_group_all_answers, groups_answers))

    print("Total count:", total_count)
    print("All answered count:", all_answered_count)
