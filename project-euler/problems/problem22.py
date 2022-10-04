"""
Problem 22
==========

Using names.txt (right click and 'Save Link/Target As...'), a 46K text file containing over
five-thousand first names, begin by sorting it into alphabetical order. Then working out the
alphabetical value for each name, multiply this value by its alphabetical position in the list
to obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN, which is worth
3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would obtain a score
of 938 × 53 = 49714.

What is the total of all the name scores in the file?
"""


def import_data(filename):
    with open(filename, "r") as name_file:
        raw_list = name_file.read()
        name_list = [name[1:-1] for name in raw_list.split(",")]
        return name_list


def name_score(name):
    score = 0

    for letter in name:
        score += ord(letter) - ord("A") + 1

    return score


def compute_list_score(name_list):
    score = 0

    for i, name in enumerate(sorted(name_list)):
        score += (i + 1) * name_score(name)

    return score


def main():
    name_list = import_data("data/p022_names.txt")
    score = compute_list_score(name_list)
    print(score)
