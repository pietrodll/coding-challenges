"""
Problem 31
==========

In the United Kingdom the currency is made up of pound (£) and pence (p). There are eight
coins in general circulation:

    1p, 2p, 5p, 10p, 20p, 50p, £1 (100p), and £2 (200p).

It is possible to make £2 in the following way:

    1×£1 + 1×50p + 2×20p + 1×5p + 1×2p + 3×1p

How many different ways can £2 be made using any number of coins?
"""


def compute_combinations(amount, coins):
    coins = sorted(coins, reverse=True)

    def aux(tot, to_use):
        combinations = []

        if tot == 0:
            combinations.append([0] * (len(coins) - to_use))

        elif to_use == len(coins) - 1 and tot % coins[-1] == 0:
            combinations.append([tot // coins[-1]])

        elif to_use < len(coins) - 1:
            for m in range(tot // coins[to_use] + 1):
                sub_combinations = aux(tot - m * coins[to_use], to_use + 1)
                combinations += [[m] + comb for comb in sub_combinations]

        return combinations

    return aux(amount, 0)


def count_possible_combinations(amount, coins):
    comb = compute_combinations(amount, coins)

    return len(comb)


def main():
    print(count_possible_combinations(200, [1, 2, 5, 10, 20, 50, 100, 200]))
