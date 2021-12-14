package day14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`

func TestParseInput(t *testing.T) {
	expectedRules := map[pair]rune{
		{'C', 'H'}: 'B',
		{'H', 'H'}: 'N',
		{'C', 'B'}: 'H',
		{'N', 'H'}: 'C',
		{'H', 'B'}: 'C',
		{'H', 'C'}: 'B',
		{'H', 'N'}: 'C',
		{'N', 'N'}: 'C',
		{'B', 'H'}: 'H',
		{'N', 'C'}: 'B',
		{'N', 'B'}: 'B',
		{'B', 'N'}: 'B',
		{'B', 'B'}: 'N',
		{'B', 'C'}: 'B',
		{'C', 'C'}: 'N',
		{'C', 'N'}: 'C',
	}

	tmp, rules := parseInput(input)
	assert.Equal(t, []rune{'N', 'N', 'C', 'B'}, tmp)
	assert.Equal(t, expectedRules, rules)
}

func TestInsertPolymers(t *testing.T) {
	tmp, rules := parseInput(input)
	expected := []string{
		"NCNBCHB",
		"NBCCNBBBCBHCB",
		"NBBBCNCCNBBNBNBBCHBHHBCHB",
		"NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB",
	}

	counter := make(map[rune]int)

	for _, expectedStep := range expected {
		tmp = insertPolymers(tmp, rules, &counter)
		assert.Equal(t, expectedStep, string(tmp))
	}
}

func TestFindMostCommonAndLeastCommon(t *testing.T) {
	counter := map[rune]int{
		'A': 5,
		'B': 6,
		'C': 1,
		'Z': 22,
	}

	mostCommon, leastCommon := findMostCommonAndLeastCommon(counter)
	assert.Equal(t, 'Z', mostCommon)
	assert.Equal(t, 'C', leastCommon)
}

func TestCountAfterSteps(t *testing.T) {
	tmpl, rules := parseInput(input)
	counter := countAfterSteps(tmpl, rules, 10)

	expected := map[rune]int{
		'B': 1749,
		'C': 298,
		'H': 161,
		'N': 865,
	}

	assert.Equal(t, expected, counter)
}

func TestMostCommonMinusLeastCommonAfterSteps(t *testing.T) {
	tmpl, rules := parseInput(input)
	assert.Equal(t, 2188189693529, mostCommonMinusLeastCommonAfterSteps(tmpl, rules, 40))
}
