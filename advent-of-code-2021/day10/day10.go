package day10

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/pietrodll/aoc2021/utils/collections"
)

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

var openToClose = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var closeToOpen = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

func findInvalidCharacters(line string) []rune {
	invalid := make([]rune, 0)
	s := collections.NewIntStack()

	for _, char := range line {
		if _, present := openToClose[char]; present {
			// it is an opening bracket
			s.Add(int(char))
		} else {
			// it is a closing bracket
			if s.IsEmpty() || closeToOpen[char] != rune(s.Pop()) {
				invalid = append(invalid, char)
			}
		}
	}

	return invalid
}

func computeSyntaxErrorScore(lines []string) int {
	scoreMapping := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	totScore := 0

	for _, line := range lines {
		invalidCharacters := findInvalidCharacters(line)

		if len(invalidCharacters) > 0 {
			totScore += scoreMapping[invalidCharacters[0]]
		}
	}

	return totScore
}

func completeLine(line string) (string, error) {
	s := collections.NewIntStack()

	for _, char := range line {
		if _, present := openToClose[char]; present {
			// it is an opening bracket
			s.Add(int(char))
		} else {
			// it is a closing bracket
			if s.IsEmpty() || closeToOpen[char] != rune(s.Pop()) {
				return "", errors.New("the line is corrupted, cannot complete")
			}
		}
	}

	completion := make([]rune, 0)

	for !s.IsEmpty() {
		completion = append(completion, openToClose[rune(s.Pop())])
	}

	return string(completion), nil
}

var completionScoreMapping = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func computeScoreFromCompletion(completion string) int {
	score := 0

	for _, char := range completion {
		score *= 5
		score += completionScoreMapping[char]
	}

	return score
}

func computeCompletionScore(lines []string) int {
	scores := make([]int, 0)

	for _, line := range lines {
		completion, err := completeLine(line)

		if err == nil {
			scores = append(scores, computeScoreFromCompletion(completion))
		}

	}

	sort.Ints(scores)

	return scores[len(scores)/2]
}

func Run(input string) {
	lines := parseInput(input)

	fmt.Println("Syntax Error Score:", computeSyntaxErrorScore(lines))
	fmt.Println("Completion Score:", computeCompletionScore(lines))
}
