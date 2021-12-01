"""
Problem 112
===========

Working from left-to-right if no digit is exceeded by the digit to its left it is called an
increasing number; for example, 134468.

Similarly if no digit is exceeded by the digit to its right it is called a decreasing number;
for example, 66420.

We shall call a positive integer that is neither increasing nor decreasing a "bouncy" number;
for example, 155349.

Clearly there cannot be any bouncy numbers below one-hundred, but just over half of the numbers
below one-thousand (525) are bouncy. In fact, the least number for which the proportion of bouncy
numbers first reaches 50% is 538.

Surprisingly, bouncy numbers become more and more common and by the time we reach 21780 the
proportion of bouncy numbers is equal to 90%.

Find the least number for which the proportion of bouncy numbers is exactly 99%.
"""


def is_increasing(n):
    x = n
    last = 10
    while x != 0:
        if x % 10 > last:
            return False
        last = x % 10
        x = x // 10
    return True


def is_decreasing(n):
    x = n
    last = 0
    while x != 0:
        if x % 10 < last:
            return False
        last = x % 10
        x = x // 10
    return True


def is_bouncy(n):
    return not is_decreasing(n) and not is_increasing(n)


def find_proportion(prop):
    proportion = 0
    n_bouncy = 0
    n = 1
    while proportion < prop:
        print(n, end='\r')
        if is_bouncy(n):
            n_bouncy += 1
        proportion = n_bouncy / n
        n += 1

    return n - 1


def run():
    print(find_proportion(0.99))

if __name__ == "__main__":
    run()
