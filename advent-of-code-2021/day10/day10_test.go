package day10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var lines = []string{
	"[({(<(())[]>[[{[]{<()<>>",
	"[(()[<>])]({[<{<<[]>>(",
	"{([(<{}[<>[]}>{[]{[(<()>",
	"(((({<>}<{<{<>}{[]{[]{}",
	"[[<[([]))<([[{}[[()]]]",
	"[{[{({}]{}}([{[{{{}}([]",
	"{<[[]]>}<{[{[{[]{()[[[]",
	"[<(<(<(<{}))><([]([]()",
	"<{([([[(<>()){}]>(<<{{",
	"<{([{{}}[<[[[<>{}]]]>[]]",
}

func TestFindInvalidCharacters(t *testing.T) {
	expectedInvalidCharacters := [][]rune{
		{},
		{},
		{'}'},
		{},
		{')'},
		{']'},
		{},
		{')'},
		{'>'},
		{},
	}

	for i, line := range lines {
		assert.Equal(t, expectedInvalidCharacters[i], findInvalidCharacters(line))
	}
}

func TestCompleteLine(t *testing.T) {
	expectedCompletions := []string{
		"}}]])})]",
		")}>]})",
		"",
		"}}>}>))))",
		"",
		"",
		"]]}}]}]}>",
		"",
		"",
		"])}>",
	}

	for i, line := range lines {
		completion, err := completeLine(line)

		if len(completion) > 0 {
			assert.Equal(t, expectedCompletions[i], completion)
			assert.Nil(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}
