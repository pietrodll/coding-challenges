"""Day 2"""

from collections import Counter, namedtuple
import re
from typing import List, Tuple

Policy = namedtuple("Policy", ["char", "min", "max"])


def parse_input(data: str) -> List[Tuple[Policy, str]]:
    pattern = re.compile(r"^(\d+)-(\d+) ([a-z]): ([a-z]+)")
    result = []

    for line in data.split("\n"):
        match = pattern.match(line)

        if not match:
            raise ValueError(f'Cannot parse "{line}"')

        policy = Policy(match.group(3), int(match.group(1)), int(match.group(2)))
        password = match.group(4)

        result.append((policy, password))

    return result


def check_password(policy: Policy, password: str):
    cnt = Counter(password)[policy.char]

    return cnt <= policy.max and cnt >= policy.min


def count_valid_passwords(policies_and_passwords: List[Tuple[Policy, str]]):
    return sum(check_password(*args) for args in policies_and_passwords)


def check_password_v2(policy: Policy, password: str):
    val = (password[policy.min - 1] == policy.char) + (
        password[policy.max - 1] == policy.char
    )
    return val == 1


def count_valid_passwords_v2(policies_and_passwords: List[Tuple[Policy, str]]):
    return sum(check_password_v2(*args) for args in policies_and_passwords)


def main(data: str):
    policies_and_passwords = parse_input(data)

    print("Valid passwords:", count_valid_passwords(policies_and_passwords))
    print(
        "Valid passwords with new policy:",
        count_valid_passwords_v2(policies_and_passwords),
    )
