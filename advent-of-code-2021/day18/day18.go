package day18

import (
	"fmt"
	"strings"
)

type SnailfishNumber struct {
	value               int
	left, right, parent *SnailfishNumber
}

func (n *SnailfishNumber) IsRegularNumber() bool {
	return n.value >= 0
}

var digits = map[rune]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

func makeNumber(value int) *SnailfishNumber {
	return &SnailfishNumber{value, nil, nil, nil}
}

func makePair(left *SnailfishNumber, right *SnailfishNumber) *SnailfishNumber {
	result := SnailfishNumber{-1, left, right, nil}
	left.parent = &result
	right.parent = &result
	return &result
}

func parseSnailfishNumberRec(data *[]rune, index *int) *SnailfishNumber {
	char := (*data)[*index]

	if val, present := digits[char]; present {
		// it is a regular number: return it and move the index forward
		*index++
		return makeNumber(val)
	}

	if char == '[' {
		// beginning of a pair
		*index++
		left := parseSnailfishNumberRec(data, index)

		if (*data)[*index] != ',' {
			panic(fmt.Errorf("invalid character at index %d", *index))
		}

		*index++
		right := parseSnailfishNumberRec(data, index)

		if (*data)[*index] != ']' {
			panic(fmt.Errorf("invalid character at index %d", *index))
		}

		*index++
		return makePair(left, right)
	}

	panic(fmt.Errorf("invalid character at index %d", *index))
}

func parseSnailfishNumber(numberStr string) *SnailfishNumber {
	data := []rune(numberStr)
	index := 0

	return parseSnailfishNumberRec(&data, &index)
}

func parseInput(input string) []*SnailfishNumber {
	lines := strings.Split(input, "\n")
	numbers := make([]*SnailfishNumber, len(lines))

	for i, line := range lines {
		numbers[i] = parseSnailfishNumber(line)
	}

	return numbers
}

func (n *SnailfishNumber) String() string {
	if n.IsRegularNumber() {
		return fmt.Sprint(n.value)
	}

	return fmt.Sprintf("[%s,%s]", n.left.String(), n.right.String())
}

func findLeftmostPairAtDepth(n *SnailfishNumber, depth int) *SnailfishNumber {
	if n.IsRegularNumber() {
		return nil
	}

	if depth == 0 {
		return n
	}

	if candidate := findLeftmostPairAtDepth(n.left, depth-1); candidate != nil {
		return candidate
	}

	return findLeftmostPairAtDepth(n.right, depth-1)
}

func findLeftmostNumberGreaterThanOrEqualTo(n *SnailfishNumber, limit int) *SnailfishNumber {
	if n.IsRegularNumber() {
		if n.value >= limit {
			return n
		}

		return nil
	}

	if candidate := findLeftmostNumberGreaterThanOrEqualTo(n.left, limit); candidate != nil {
		return candidate
	}

	return findLeftmostNumberGreaterThanOrEqualTo(n.right, limit)
}

func findClosestNumberLeft(n *SnailfishNumber) *SnailfishNumber {
	curr := n

	for curr.parent != nil && curr != curr.parent.right {
		curr = curr.parent
	}

	if curr.parent == nil {
		return nil
	}

	curr = curr.parent.left

	for !curr.IsRegularNumber() {
		curr = curr.right
	}

	return curr
}

func findClosestNumberRight(n *SnailfishNumber) *SnailfishNumber {
	curr := n

	for curr.parent != nil && curr != curr.parent.left {
		curr = curr.parent
	}

	if curr.parent == nil {
		return nil
	}

	curr = curr.parent.right

	for !curr.IsRegularNumber() {
		curr = curr.left
	}

	return curr
}

func explodeStep(n *SnailfishNumber) bool {
	// check if a pair must explode
	exploding := findLeftmostPairAtDepth(n, 4)

	if exploding == nil {
		return false
	}

	if closestLeft := findClosestNumberLeft(exploding); closestLeft != nil {
		closestLeft.value += exploding.left.value
	}

	if closestRight := findClosestNumberRight(exploding); closestRight != nil {
		closestRight.value += exploding.right.value
	}

	// substitute the exploding pair by a regular number 0
	exploding.value = 0
	exploding.left = nil
	exploding.right = nil

	return true
}

func splitStep(n *SnailfishNumber) bool {
	toSplit := findLeftmostNumberGreaterThanOrEqualTo(n, 10)

	if toSplit == nil {
		return false
	}

	val := toSplit.value

	toSplit.value = -1
	toSplit.left = &SnailfishNumber{val / 2, nil, nil, toSplit}
	toSplit.right = &SnailfishNumber{val - toSplit.left.value, nil, nil, toSplit}

	return true
}

// reduces the number, returning false if the number was left unchanged
func reduceStep(n *SnailfishNumber) bool {
	if hasExploded := explodeStep(n); hasExploded {
		return true
	}

	if hasSplitted := splitStep(n); hasSplitted {
		return true
	}

	return false
}

func reduce(n *SnailfishNumber) {
	for reduceStep(n) {
	}
}

func sum(a, b *SnailfishNumber) *SnailfishNumber {
	s := makePair(a, b)
	reduce(s)
	return s
}

func sumAll(numbers []*SnailfishNumber) *SnailfishNumber {
	tot := numbers[0]

	for _, number := range numbers[1:] {
		tot = sum(tot, number)
	}

	return tot
}

func (n *SnailfishNumber) Magnitude() int {
	if n.IsRegularNumber() {
		return n.value
	}

	return 3*n.left.Magnitude() + 2*n.right.Magnitude()
}

func findHighestMagnitudeSum(numbersStr []string) *SnailfishNumber {
	var (
		max    int
		result *SnailfishNumber
	)

	for i, aStr := range numbersStr {
		for j, bStr := range numbersStr {
			if i != j {
				a := parseSnailfishNumber(aStr)
				b := parseSnailfishNumber(bStr)
				s := sum(a, b)

				if m := s.Magnitude(); m > max {
					result = s
					max = m
				}
			}
		}
	}

	return result
}

func Run(input string) {
	numbers := parseInput(input)
	s := sumAll(numbers)

	fmt.Println("Magnitude of the sum:", s.Magnitude())
	fmt.Println(
		"Highest magnitude possible by summing two numbers:",
		findHighestMagnitudeSum(strings.Split(input, "\n")).Magnitude(),
	)
}
