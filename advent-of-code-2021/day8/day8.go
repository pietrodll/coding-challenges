package day8

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Entry struct {
	pattern []string
	digits  []string
}

func sortString(str string) string {
	chars := strings.Split(str, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

func parseInput(input string) []Entry {
	lines := strings.Split(input, "\n")
	entries := make([]Entry, len(lines))

	for i, line := range lines {
		split := strings.Split(line, " | ")
		pattern := strings.Split(split[0], " ")
		digits := strings.Split(split[1], " ")

		for i, segments := range pattern {
			pattern[i] = sortString(segments)
		}

		for i, digit := range digits {
			digits[i] = sortString(digit)
		}

		entries[i] = Entry{pattern, digits}
	}

	return entries
}

func countUniqueSegmentDigits(entries []Entry) int {
	cnt := 0

	for _, entry := range entries {
		for _, digit := range entry.digits {
			if len(digit) == 2 || len(digit) == 3 || len(digit) == 4 || len(digit) == 7 {
				cnt += 1
			}
		}
	}

	return cnt
}

func readDigits(segmentsToDigit map[string]rune, segmentDigits []string) int {
	runes := make([]rune, len(segmentDigits))

	for i, segments := range segmentDigits {
		runes[i] = segmentsToDigit[segments]
	}

	value, err := strconv.Atoi(string(runes))

	if err != nil {
		panic(err)
	}

	return value
}

// the characters in the strings must be in alphabetical order
func intersectionLength(a string, b string) int {
	i, j := 0, 0
	runesA := []rune(a)
	runesB := []rune(b)
	intersectionLength := 0

	for i < len(runesA) && j < len(runesB) {
		if runesA[i] == runesB[j] {
			intersectionLength++
			i++
			j++
		} else if runesA[i] < runesB[j] {
			i++
		} else {
			j++
		}
	}

	return intersectionLength
}

func decodeDigits(entry Entry) int {
	segmentsToDigit := make(map[string]rune)
	digitToSegments := make(map[rune]string)

	byLength := make(map[int][]string)

	for _, segments := range entry.pattern {
		byLength[len(segments)] = append(byLength[len(segments)], segments)
	}

	// the only digit made of 2 segments is 1
	segmentsToDigit[byLength[2][0]] = '1'
	digitToSegments['1'] = byLength[2][0]

	// the only digit made of 3 segments is 7
	segmentsToDigit[byLength[3][0]] = '7'
	digitToSegments['7'] = byLength[3][0]

	// the only digit made of 4 segments is 4
	segmentsToDigit[byLength[4][0]] = '4'
	digitToSegments['4'] = byLength[4][0]

	// the only digit made of 7 segments is 8
	segmentsToDigit[byLength[7][0]] = '8'
	digitToSegments['8'] = byLength[7][0]

	// 2, 3, and 5 are the digits made of 5 segments
	for _, segments := range byLength[5] {
		if intersectionLength(segments, digitToSegments['4']) == 2 {
			// among them, 2 is the only one with 2 segments in common with 4
			segmentsToDigit[segments] = '2'
			digitToSegments['2'] = segments
		} else if intersectionLength(segments, digitToSegments['1']) == 2 {
			// among them, 3 is the only one with 2 segments in common with 1
			segmentsToDigit[segments] = '3'
			digitToSegments['3'] = segments
		} else {
			// the remaining one is 5
			segmentsToDigit[segments] = '5'
			digitToSegments['5'] = segments
		}
	}

	// 0, 6 and 9 are the digits made of 6 segments
	for _, segments := range byLength[6] {
		if intersectionLength(segments, digitToSegments['1']) == 1 {
			// among them, 6 is the only one with 1 segment in common with 1
			segmentsToDigit[segments] = '6'
			digitToSegments['6'] = segments
		} else if intersectionLength(segments, digitToSegments['4']) == 4 {
			// among them, 9 is the only one with 4 segments in common with 4
			segmentsToDigit[segments] = '9'
			digitToSegments['9'] = segments
		} else {
			// the remaining one is 0
			segmentsToDigit[segments] = '0'
			digitToSegments['0'] = segments
		}
	}

	return readDigits(segmentsToDigit, entry.digits)
}

func decodeAndComputeSum(entries []Entry) int {
	tot := 0

	for _, entry := range entries {
		tot += decodeDigits(entry)
	}

	return tot
}

func Run(input string) {
	entries := parseInput(input)

	fmt.Println("Number of unique segment digits:", countUniqueSegmentDigits(entries))
	fmt.Println("Sum of decoded digits:", decodeAndComputeSum(entries))
}
