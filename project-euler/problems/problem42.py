"""
Problem 42
==========

The nth term of the sequence of triangle numbers is given by, tn = ½n(n+1); so the first ten
triangle numbers are:

1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...

By converting each letter in a word to a number corresponding to its alphabetical position and
adding these values we form a word value.
For example, the word value for SKY is 19 + 11 + 25 = 55 = t10.
If the word value is a triangle number then we shall call the word a triangle word.

Using words.txt (right click and 'Save Link/Target As...'), a 16K text file containing nearly
two-thousand common English words, how many are triangle words?
"""

# n² + n - 2x = 0
# delta = 1 + 8x
# n = (-1 +- sqrt(1 + 8x)) / 2

from math import sqrt


def is_triangle_num(x):
    root = sqrt(1 + 8 * x)

    return int(root) == root and root % 2 == 1


def word_value(word):
    word = word.upper()
    return sum(ord(c) - ord("A") + 1 for c in word)


def is_triangle_word(word: str):
    return is_triangle_num(word_value(word))


def read_words(filename):
    with open(filename) as word_file:
        words_with_quotes = word_file.read().split(",")
        return [word[1:-1] for word in words_with_quotes]


def count_triangle_words(filename):
    words = read_words(filename)

    return sum(map(is_triangle_word, words))


def main():
    assert is_triangle_num(36) and is_triangle_num(55)

    print(count_triangle_words("data/p042_words.txt"))
