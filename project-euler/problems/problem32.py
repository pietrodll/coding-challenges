"""
Problem 32
==========

We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once; for example, the 5-digit number, 15234, is 1 through 5 pandigital.

The product 7254 is unusual, as the identity, 39 x 186 = 7254, containing multiplicand, multiplier, and product is 1 through 9 pandigital.

Find the sum of all products whose multiplicand/multiplier/product identity can be written as a 1 through 9 pandigital.

HINT: Some products can be obtained in more than one way so be sure to only include it once in your sum.
"""

# A number with N digits is between 10^(N-1) and 10^N
# If the multiplicand has N digits and the mulplier has M, the product P verifies:
# 10^(N-1) * 10^(M-1) < P < 10^N * 10^M => N + M - 2 < digits of P < N + M
# But digits of P = 9 - N - M
# => N + M - 2 < 9 - M - N    &    9 - N - M < N + M
# => 9 < 2(N + M) < 11
# => 2(N + M) = 10
# => N + M = 5
# => 3 < digits of P < 5 => P has 4 digits

from typing import Generator, List, Set, Tuple

from utils.digits import digits_to_num


def generate_combinations(
    digits: Set[int], size: int
) -> Generator[List[int], None, None]:
    assert size <= len(digits)

    if size == 0:
        yield []
        return

    for first_digit in digits:
        for rest in generate_combinations(digits - {first_digit}, size - 1):
            yield [first_digit] + rest


def generate_numbers_with_digits(
    num_digits: int,
) -> Generator[Tuple[int, Set[int]], None, None]:
    for comb in generate_combinations(set(range(1, 10)), num_digits):
        yield digits_to_num(comb, reverse=True), set(comb)


def iter_potential_products_with_digits():
    """Generates all possible products, i.e. all numbers with 4 distinct digits"""

    return generate_numbers_with_digits(4)


def is_product_of_numbers_with_digits(product: int, available_digits: Set[int]) -> bool:
    """Returns true if the given number is a pandigital product, i.e. if we can find two
    numbers x and y such that x * y = P and the digits of x, y, and P cover all digits from 1 to 9.
    """
    for size in (1, 2, 3):
        for first_digits in generate_combinations(available_digits, size):
            first = digits_to_num(first_digits, reverse=True)

            if product % first == 0:
                second = product // first
                possible_second_digits = available_digits - set(first_digits)
                second_size = len(available_digits) - size

                if any(
                    second == digits_to_num(second_digits, reverse=True)
                    for second_digits in generate_combinations(
                        possible_second_digits, second_size
                    )
                ):
                    return True

    return False


def find_all_products():
    """Returns the list of all pandigital products"""
    products = []
    all_digits = set(range(1, 10))

    for product, digits in iter_potential_products_with_digits():
        remaining_digits = all_digits - digits
        if is_product_of_numbers_with_digits(product, remaining_digits):
            products.append(product)

    return products


def main():
    print(sum(find_all_products()))


if __name__ == "__main__":
    main()
