package day3

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

func TestParseInput(t *testing.T) {
	values, bitNum := parseInput(input)

	assert.Equal(t, 12, len(values))
	assert.Equal(t, 5, bitNum)
}

func TestGetBitByIndex(t *testing.T) {
	binary, err := strconv.ParseUint("10110", 2, 64)

	assert.Nil(t, err)

	assert.Equal(t, uint64(22), binary)

	assert.Equal(t, uint64(0), getBitByIndex(binary, 0))
	assert.Equal(t, uint64(1), getBitByIndex(binary, 1))
	assert.Equal(t, uint64(1), getBitByIndex(binary, 2))
}

func TestFindMostCommonBitByIndex(t *testing.T) {
	values, _ := parseInput(input)

	assert.Equal(t, uint64(0), findMostCommonBitByIndex(values, 0))
	assert.Equal(t, uint64(1), findMostCommonBitByIndex(values, 1))
	assert.Equal(t, uint64(1), findMostCommonBitByIndex(values, 2))
	assert.Equal(t, uint64(0), findMostCommonBitByIndex(values, 3))
	assert.Equal(t, uint64(1), findMostCommonBitByIndex(values, 4))

	// if there is the same number of 0s and 1s, we return 1 as the most common bit.
	values = []uint64{22, 23}
	assert.Equal(t, uint64(1), findMostCommonBitByIndex(values, 0))
}

func TestFindEpsilonAndGamma(t *testing.T) {
	epsilon, gamma := findEpsilonAndGamma(parseInput(input))

	assert.Equal(t, 9, epsilon)
	assert.Equal(t, 22, gamma)
}

func TestGetConsumption(t *testing.T) {
	assert.Equal(t, 198, getConsumption(parseInput(input)))
}

func TestGetLifeSupportRating(t *testing.T) {
	values, bitNum := parseInput(input)

	assert.Equal(t, 23, findOxygen(values, bitNum))
	assert.Equal(t, 10, findCo2(values, bitNum))
	assert.Equal(t, 230, getLifeSupportRating(values, bitNum))
}
