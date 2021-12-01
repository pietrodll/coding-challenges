"""
Problem 39
==========


If p is the perimeter of a right angle triangle with integral length sides, {a,b,c}, there
are exactly three solutions for p = 120.

{20,48,52}, {24,45,51}, {30,40,50}

For which value of p ≤ 1000, is the number of solutions maximised?
"""

def indexmax(L):
    imax = 0
    m = L[0]

    for i, x in enumerate(L):
        if x > m:
            imax = i
            m = x

    return imax


def main():
    solutions = [0] * 1001

    for h in range(1, 1000):
        for c1 in range(1, h):
            for c2 in range(1, h):
                p = h + c1 + c2

                if p > 1000:
                    break

                if h * h == c1 * c1 + c2 * c2:
                    solutions[p] += 1

    print(indexmax(solutions))


if __name__ == "__main__":
    main()
