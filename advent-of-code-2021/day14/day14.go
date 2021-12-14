package day14

import (
	"fmt"
	"strings"
)

type pair struct {
	left  rune
	right rune
}

func parseInput(input string) ([]rune, map[pair]rune) {
	blocks := strings.Split(input, "\n\n")

	rules := make(map[pair]rune)

	for _, rule := range strings.Split(blocks[1], "\n") {
		runes := []rune(rule)
		rules[pair{runes[0], runes[1]}] = runes[len(runes)-1]
	}

	return []rune(blocks[0]), rules
}

func insertPolymers(current []rune, rules map[pair]rune, counter *map[rune]int) []rune {
	result := make([]rune, 0, 2*len(current)-1)
	result = append(result, current[0])

	for i := 1; i < len(current); i++ {
		if toInsert, present := rules[pair{current[i-1], current[i]}]; present {
			(*counter)[toInsert]++
			result = append(result, toInsert, current[i])
		} else {
			result = append(result, current[i])
		}
	}

	return result
}

func findMostCommonAndLeastCommon(counter map[rune]int) (rune, rune) {
	var (
		mostCommon, leastCommon           rune
		mostCommonCount, leastCommonCount int
	)

	for char, count := range counter {
		mostCommon = char
		leastCommon = char
		mostCommonCount = count
		leastCommonCount = count
		break
	}

	for char, count := range counter {
		if count > mostCommonCount {
			mostCommon = char
			mostCommonCount = count
		}

		if count < leastCommonCount {
			leastCommon = char
			leastCommonCount = count
		}
	}

	return mostCommon, leastCommon
}

func countPairs(tmpl []rune) map[pair]int {
	counter := make(map[pair]int)

	for i := 1; i < len(tmpl); i++ {
		counter[pair{tmpl[i-1], tmpl[i]}]++
	}

	return counter
}

func countAfterSteps(tmpl []rune, rules map[pair]rune, steps int) map[rune]int {
	// we store the template as a counter of pairs: for each step, each pair spawns
	// new pairs depending on its count

	// initialize the counter
	counter := make(map[rune]int)

	for _, char := range tmpl {
		counter[char]++
	}

	// initialize the pair counter
	pairs := countPairs(tmpl)

	for step := 0; step < steps; step++ {
		next := make(map[pair]int)

		for p, count := range pairs {
			if toInsert, present := rules[p]; present {
				counter[toInsert] += count
				next[pair{p.left, toInsert}] += count
				next[pair{toInsert, p.right}] += count
			}
		}

		pairs = next
	}

	return counter
}

func mostCommonMinusLeastCommonAfterSteps(tmpl []rune, rules map[pair]rune, steps int) int {
	counter := countAfterSteps(tmpl, rules, steps)
	mostCommon, leastCommon := findMostCommonAndLeastCommon(counter)

	return counter[mostCommon] - counter[leastCommon]
}

func Run(input string) {
	tmpl, rules := parseInput(input)

	fmt.Println(
		"Most common - least common after 10 steps:",
		mostCommonMinusLeastCommonAfterSteps(tmpl, rules, 10),
	)
	fmt.Println(
		"Most common - least common after 40 steps:",
		mostCommonMinusLeastCommonAfterSteps(tmpl, rules, 40),
	)
}
