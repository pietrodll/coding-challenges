"""
Problem 17
==========

If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there
are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how
many letters would be used?

NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23
letters and 115 (one hundred and fifteen) contains 20 letters. The use of "and" when writing out
numbers is in compliance with British usage.
"""

import re


ONES = ['', 'one', 'two', 'three', 'four', 'five', 'six', 'seven', 'eight', 'nine', 'ten',
        'eleven', 'twelve', 'thirteen', 'fourteen', 'fifteen', 'sixteen', 'seventeen',
        'eighteen', 'nineteen']
TENS = ['', 'ten', 'twenty', 'thirty', 'forty', 'fifty', 'sixty', 'seventy', 'eighty', 'ninety']
HUNDRED = 'hundred'
THOUSAND = 'thousand'


def write_in_letters(n):
    text = ''
    m = n % 100

    if m < 20:
        text = ONES[m]

    else:
        text = TENS[m // 10]

        if m % 10 > 0:
            text += f'-{ONES[m % 10]}'

    n //= 100

    if n > 0 and n % 10 > 0:
        if text != '':
            text = f'{ONES[n % 10]} {HUNDRED} and {text}'

        else:
            text = f'{ONES[n % 10]} {HUNDRED}'

    n //= 10

    if n > 0:
        text = f'{ONES[n]} {THOUSAND} {text}'

    return text


def count_letters(n):
    text = write_in_letters(n)
    return len(re.sub(r'[\s\-]', '', text))


def sum_count_letters(start, end):
    return sum(map(count_letters, range(start, end + 1)), 0)


if __name__ == "__main__":
    assert count_letters(342) == 23
    assert count_letters(115) == 20
    assert sum_count_letters(1, 5) == 19
    print(sum_count_letters(1, 1000))
