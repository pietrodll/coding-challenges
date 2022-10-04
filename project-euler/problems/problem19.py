"""
Problem 19
==========

You are given the following information, but you may prefer to do some research for yourself.

    1 Jan 1900 was a Monday.

    Thirty days has September,
    April, June and November.
    All the rest have thirty-one,
    Saving February alone,
    Which has twenty-eight, rain or shine.
    And on leap years, twenty-nine.

    A leap year occurs on any year evenly divisible by 4, but not on a century unless it is
    divisible by 400.

How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to
31 Dec 2000)?
"""

DAYS_OF_MONTH = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]


def is_leap(year):
    return year % 400 == 0 or (year % 4 == 0 and year % 100 != 0)


def get_month_days(year):
    res = DAYS_OF_MONTH[:]

    if is_leap(year):
        res[1] += 1

    return res


def count_sundays_first_in_year(year, day_of_jan_first):
    days = get_month_days(year)
    first = day_of_jan_first
    count = 0

    for month_days in days:
        if first == 0:
            count += 1

        first = (first + month_days) % 7

    return count, first


def count_sundays_first_in_years(start, end, day_of_jan_first):
    total = 0
    first = day_of_jan_first

    for year in range(start, end + 1):
        count, first = count_sundays_first_in_year(year, first)
        total += count

    return total


def weekday_delay(base_year, weekday, target_year):
    delay = sum(
        (366 if is_leap(year) else 365) for year in range(base_year, target_year)
    )

    return (weekday + delay) % 7


def sunday_count_19th_cent():
    first_day_1901 = weekday_delay(1900, 1, 1901)

    return count_sundays_first_in_years(1901, 2000, first_day_1901)


def main():
    assert not is_leap(1900)
    assert is_leap(2000)

    assert count_sundays_first_in_year(2020, 3) == (2, 5)
    assert count_sundays_first_in_years(2020, 2020, 3) == 2

    assert count_sundays_first_in_year(2009, 4) == (3, 5)

    print(sunday_count_19th_cent())
