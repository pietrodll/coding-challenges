package day9

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `2199943210
3987894921
9856789892
8767896789
9899965678`

func TestParseInput(t *testing.T) {
	expected := Grid{
		[][]int{{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
			{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
			{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
			{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
			{9, 8, 9, 9, 9, 6, 5, 6, 7, 8}},
		5,
		10,
	}

	assert.Equal(t, expected, parseInput(input))
}

func TestFindLowPoints(t *testing.T) {
	grid := parseInput(input)
	expected := []GridPoint{{0, 1}, {0, 9}, {2, 2}, {4, 6}}

	assert.Equal(t, expected, grid.findLowPoints())
	assert.Equal(t, 15, grid.totalRiskLevel())
}

func TestFindBasins(t *testing.T) {
	grid := parseInput(input)
	basins := grid.findBasins()

	assert.Len(t, basins, 4)

	assert.ElementsMatch(t, []GridPoint{{0, 0}, {1, 0}, {0, 1}}, basins[0])
	assert.ElementsMatch(t, []GridPoint{{0, 9}, {0, 8}, {0, 7}, {0, 6}, {0, 5}, {1, 9}, {2, 9}, {1, 8}, {1, 6}}, basins[1])
	assert.ElementsMatch(t, []GridPoint{{1, 2}, {1, 3}, {1, 4}, {2, 1}, {2, 2}, {2, 3}, {2, 4}, {2, 5}, {3, 0}, {3, 1}, {3, 2}, {3, 3}, {3, 4}, {4, 1}}, basins[2])
	assert.ElementsMatch(t, []GridPoint{{2, 7}, {3, 6}, {3, 7}, {3, 8}, {4, 5}, {4, 6}, {4, 7}, {4, 8}, {4, 9}}, basins[3])
}
