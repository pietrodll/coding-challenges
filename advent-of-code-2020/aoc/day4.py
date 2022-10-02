"""Day 4"""

import re
from typing import List, Dict


def parse_input(data: str) -> List[Dict[str, str]]:
    blocks = data.split("\n\n")
    passports = []

    pattern = re.compile(r"^([a-z]{3}):(\S+)$")

    for block in blocks:
        passport = {}
        fields = re.split(r"\s+", block)

        for field in fields:
            match = pattern.match(field)

            if not match:
                raise ValueError(f"Cannot parse {field}")

            passport[match.group(1)] = match.group(2)

        passports.append(passport)

    return passports


def check_passport(passport: Dict[str, str]):
    return len(passport) == 8 or (len(passport) == 7 and "cid" not in passport)


def validate_birth_year(byr: str):
    if len(byr) != 4:
        return False

    val = int(byr)

    return val >= 1920 and val <= 2002


def validate_issue_year(iyr: str):
    if len(iyr) != 4:
        return False

    val = int(iyr)

    return val >= 2010 and val <= 2020


def validate_expiration_year(eyr: str):
    if len(eyr) != 4:
        return False

    val = int(eyr)

    return val >= 2020 and val <= 2030


def validate_height(hgt: str):
    match = re.match(r"^(\d+)(cm|in)$", hgt)

    if not match:
        return False

    val = int(match.group(1))

    if match.group(2) == "in":
        return val >= 59 and val <= 76

    return val >= 150 and val <= 193


def validate_hair_color(hcl: str):
    return bool(re.match(r"^#[a-z0-9]{6}$", hcl))


ALLOWED_EYE_COLORS = {"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}


def validate_eye_color(ecl: str):
    return ecl in ALLOWED_EYE_COLORS


def validate_passport_id(pid: str):
    return bool(re.fullmatch(r"\d{9}", pid))


def validate_country_id(_: str):
    return True


VALIDATION_SCHEMA = {
    "byr": validate_birth_year,
    "iyr": validate_issue_year,
    "eyr": validate_expiration_year,
    "hgt": validate_height,
    "hcl": validate_hair_color,
    "ecl": validate_eye_color,
    "pid": validate_passport_id,
    "cid": validate_country_id,
}


def validate_passport(passport: Dict[str, str]):
    return all(
        validation(passport.get(key, ""))
        for (key, validation) in VALIDATION_SCHEMA.items()
    )


def main(data: str):
    passports = parse_input(data)

    valid_cnt = sum(map(check_passport, passports))
    really_valid_cnt = sum(map(validate_passport, passports))

    print("Valid passports:", valid_cnt)
    print("Really valid passports:", really_valid_cnt)
