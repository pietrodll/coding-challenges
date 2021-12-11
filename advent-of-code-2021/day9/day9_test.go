package day9

import (
	"testing"

	"github.com/pietrodll/aoc2021/utils/grid"
	"github.com/stretchr/testify/assert"
)

var input = `2199943210
3987894921
9856789892
8767896789
9899965678`

func TestParseInput(t *testing.T) {
	expected := grid.NewGrid(
		[][]int{{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
			{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
			{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
			{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
			{9, 8, 9, 9, 9, 6, 5, 6, 7, 8}})

	assert.Equal(t, expected, parseInput(input))
}

func TestFindLowPoints(t *testing.T) {
	g := parseInput(input)
	expected := []grid.GridPoint{{I: 0, J: 1}, {I: 0, J: 9}, {I: 2, J: 2}, {I: 4, J: 6}}

	assert.Equal(t, expected, findLowPoints(&g))
	assert.Equal(t, 15, totalRiskLevel(&g))
}

func TestFindBasins(t *testing.T) {
	g := parseInput(input)
	basins := findBasins(&g)

	assert.Len(t, basins, 4)

	assert.ElementsMatch(t, []grid.GridPoint{{I: 0, J: 0}, {I: 1, J: 0}, {I: 0, J: 1}}, basins[0])
	assert.ElementsMatch(t, []grid.GridPoint{{I: 0, J: 9}, {I: 0, J: 8}, {I: 0, J: 7}, {I: 0, J: 6}, {I: 0, J: 5}, {I: 1, J: 9}, {I: 2, J: 9}, {I: 1, J: 8}, {I: 1, J: 6}}, basins[1])
	assert.ElementsMatch(t, []grid.GridPoint{{I: 1, J: 2}, {I: 1, J: 3}, {I: 1, J: 4}, {I: 2, J: 1}, {I: 2, J: 2}, {I: 2, J: 3}, {I: 2, J: 4}, {I: 2, J: 5}, {I: 3, J: 0}, {I: 3, J: 1}, {I: 3, J: 2}, {I: 3, J: 3}, {I: 3, J: 4}, {I: 4, J: 1}}, basins[2])
	assert.ElementsMatch(t, []grid.GridPoint{{I: 2, J: 7}, {I: 3, J: 6}, {I: 3, J: 7}, {I: 3, J: 8}, {I: 4, J: 5}, {I: 4, J: 6}, {I: 4, J: 7}, {I: 4, J: 8}, {I: 4, J: 9}}, basins[3])
}

func TestFindAndMultiplyThreeLargestBasins(t *testing.T) {
	g := parseInput(input)
	assert.Equal(t, 1134, findAndMultiplyThreeLargestBasins(&g))
}
