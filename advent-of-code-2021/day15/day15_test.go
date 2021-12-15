package day15

import (
	"testing"

	"github.com/pietrodll/aoc2021/utils/grid"
	"github.com/stretchr/testify/assert"
)

var input = `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`

func TestParseInput(t *testing.T) {
	expected := grid.NewGrid([][]int{
		{1, 1, 6, 3, 7, 5, 1, 7, 4, 2},
		{1, 3, 8, 1, 3, 7, 3, 6, 7, 2},
		{2, 1, 3, 6, 5, 1, 1, 3, 2, 8},
		{3, 6, 9, 4, 9, 3, 1, 5, 6, 9},
		{7, 4, 6, 3, 4, 1, 7, 1, 1, 1},
		{1, 3, 1, 9, 1, 2, 8, 1, 3, 7},
		{1, 3, 5, 9, 9, 1, 2, 4, 2, 1},
		{3, 1, 2, 5, 4, 2, 1, 6, 3, 9},
		{1, 2, 9, 3, 1, 3, 8, 5, 2, 1},
		{2, 3, 1, 1, 9, 4, 4, 5, 8, 1},
	})

	assert.Equal(t, expected, parseInput(input))
}

func TestGenerateWholeMap(t *testing.T) {
	expected := grid.NewGrid([][]int{
		{1, 1, 6, 3, 7, 5, 1, 7, 4, 2, 2, 2, 7, 4, 8, 6, 2, 8, 5, 3, 3, 3, 8, 5, 9, 7, 3, 9, 6, 4, 4, 4, 9, 6, 1, 8, 4, 1, 7, 5, 5, 5, 1, 7, 2, 9, 5, 2, 8, 6},
		{1, 3, 8, 1, 3, 7, 3, 6, 7, 2, 2, 4, 9, 2, 4, 8, 4, 7, 8, 3, 3, 5, 1, 3, 5, 9, 5, 8, 9, 4, 4, 6, 2, 4, 6, 1, 6, 9, 1, 5, 5, 7, 3, 5, 7, 2, 7, 1, 2, 6},
		{2, 1, 3, 6, 5, 1, 1, 3, 2, 8, 3, 2, 4, 7, 6, 2, 2, 4, 3, 9, 4, 3, 5, 8, 7, 3, 3, 5, 4, 1, 5, 4, 6, 9, 8, 4, 4, 6, 5, 2, 6, 5, 7, 1, 9, 5, 5, 7, 6, 3},
		{3, 6, 9, 4, 9, 3, 1, 5, 6, 9, 4, 7, 1, 5, 1, 4, 2, 6, 7, 1, 5, 8, 2, 6, 2, 5, 3, 7, 8, 2, 6, 9, 3, 7, 3, 6, 4, 8, 9, 3, 7, 1, 4, 8, 4, 7, 5, 9, 1, 4},
		{7, 4, 6, 3, 4, 1, 7, 1, 1, 1, 8, 5, 7, 4, 5, 2, 8, 2, 2, 2, 9, 6, 8, 5, 6, 3, 9, 3, 3, 3, 1, 7, 9, 6, 7, 4, 1, 4, 4, 4, 2, 8, 1, 7, 8, 5, 2, 5, 5, 5},
		{1, 3, 1, 9, 1, 2, 8, 1, 3, 7, 2, 4, 2, 1, 2, 3, 9, 2, 4, 8, 3, 5, 3, 2, 3, 4, 1, 3, 5, 9, 4, 6, 4, 3, 4, 5, 2, 4, 6, 1, 5, 7, 5, 4, 5, 6, 3, 5, 7, 2},
		{1, 3, 5, 9, 9, 1, 2, 4, 2, 1, 2, 4, 6, 1, 1, 2, 3, 5, 3, 2, 3, 5, 7, 2, 2, 3, 4, 6, 4, 3, 4, 6, 8, 3, 3, 4, 5, 7, 5, 4, 5, 7, 9, 4, 4, 5, 6, 8, 6, 5},
		{3, 1, 2, 5, 4, 2, 1, 6, 3, 9, 4, 2, 3, 6, 5, 3, 2, 7, 4, 1, 5, 3, 4, 7, 6, 4, 3, 8, 5, 2, 6, 4, 5, 8, 7, 5, 4, 9, 6, 3, 7, 5, 6, 9, 8, 6, 5, 1, 7, 4},
		{1, 2, 9, 3, 1, 3, 8, 5, 2, 1, 2, 3, 1, 4, 2, 4, 9, 6, 3, 2, 3, 4, 2, 5, 3, 5, 1, 7, 4, 3, 4, 5, 3, 6, 4, 6, 2, 8, 5, 4, 5, 6, 4, 7, 5, 7, 3, 9, 6, 5},
		{2, 3, 1, 1, 9, 4, 4, 5, 8, 1, 3, 4, 2, 2, 1, 5, 5, 6, 9, 2, 4, 5, 3, 3, 2, 6, 6, 7, 1, 3, 5, 6, 4, 4, 3, 7, 7, 8, 2, 4, 6, 7, 5, 5, 4, 8, 8, 9, 3, 5},
		{2, 2, 7, 4, 8, 6, 2, 8, 5, 3, 3, 3, 8, 5, 9, 7, 3, 9, 6, 4, 4, 4, 9, 6, 1, 8, 4, 1, 7, 5, 5, 5, 1, 7, 2, 9, 5, 2, 8, 6, 6, 6, 2, 8, 3, 1, 6, 3, 9, 7},
		{2, 4, 9, 2, 4, 8, 4, 7, 8, 3, 3, 5, 1, 3, 5, 9, 5, 8, 9, 4, 4, 6, 2, 4, 6, 1, 6, 9, 1, 5, 5, 7, 3, 5, 7, 2, 7, 1, 2, 6, 6, 8, 4, 6, 8, 3, 8, 2, 3, 7},
		{3, 2, 4, 7, 6, 2, 2, 4, 3, 9, 4, 3, 5, 8, 7, 3, 3, 5, 4, 1, 5, 4, 6, 9, 8, 4, 4, 6, 5, 2, 6, 5, 7, 1, 9, 5, 5, 7, 6, 3, 7, 6, 8, 2, 1, 6, 6, 8, 7, 4},
		{4, 7, 1, 5, 1, 4, 2, 6, 7, 1, 5, 8, 2, 6, 2, 5, 3, 7, 8, 2, 6, 9, 3, 7, 3, 6, 4, 8, 9, 3, 7, 1, 4, 8, 4, 7, 5, 9, 1, 4, 8, 2, 5, 9, 5, 8, 6, 1, 2, 5},
		{8, 5, 7, 4, 5, 2, 8, 2, 2, 2, 9, 6, 8, 5, 6, 3, 9, 3, 3, 3, 1, 7, 9, 6, 7, 4, 1, 4, 4, 4, 2, 8, 1, 7, 8, 5, 2, 5, 5, 5, 3, 9, 2, 8, 9, 6, 3, 6, 6, 6},
		{2, 4, 2, 1, 2, 3, 9, 2, 4, 8, 3, 5, 3, 2, 3, 4, 1, 3, 5, 9, 4, 6, 4, 3, 4, 5, 2, 4, 6, 1, 5, 7, 5, 4, 5, 6, 3, 5, 7, 2, 6, 8, 6, 5, 6, 7, 4, 6, 8, 3},
		{2, 4, 6, 1, 1, 2, 3, 5, 3, 2, 3, 5, 7, 2, 2, 3, 4, 6, 4, 3, 4, 6, 8, 3, 3, 4, 5, 7, 5, 4, 5, 7, 9, 4, 4, 5, 6, 8, 6, 5, 6, 8, 1, 5, 5, 6, 7, 9, 7, 6},
		{4, 2, 3, 6, 5, 3, 2, 7, 4, 1, 5, 3, 4, 7, 6, 4, 3, 8, 5, 2, 6, 4, 5, 8, 7, 5, 4, 9, 6, 3, 7, 5, 6, 9, 8, 6, 5, 1, 7, 4, 8, 6, 7, 1, 9, 7, 6, 2, 8, 5},
		{2, 3, 1, 4, 2, 4, 9, 6, 3, 2, 3, 4, 2, 5, 3, 5, 1, 7, 4, 3, 4, 5, 3, 6, 4, 6, 2, 8, 5, 4, 5, 6, 4, 7, 5, 7, 3, 9, 6, 5, 6, 7, 5, 8, 6, 8, 4, 1, 7, 6},
		{3, 4, 2, 2, 1, 5, 5, 6, 9, 2, 4, 5, 3, 3, 2, 6, 6, 7, 1, 3, 5, 6, 4, 4, 3, 7, 7, 8, 2, 4, 6, 7, 5, 5, 4, 8, 8, 9, 3, 5, 7, 8, 6, 6, 5, 9, 9, 1, 4, 6},
		{3, 3, 8, 5, 9, 7, 3, 9, 6, 4, 4, 4, 9, 6, 1, 8, 4, 1, 7, 5, 5, 5, 1, 7, 2, 9, 5, 2, 8, 6, 6, 6, 2, 8, 3, 1, 6, 3, 9, 7, 7, 7, 3, 9, 4, 2, 7, 4, 1, 8},
		{3, 5, 1, 3, 5, 9, 5, 8, 9, 4, 4, 6, 2, 4, 6, 1, 6, 9, 1, 5, 5, 7, 3, 5, 7, 2, 7, 1, 2, 6, 6, 8, 4, 6, 8, 3, 8, 2, 3, 7, 7, 9, 5, 7, 9, 4, 9, 3, 4, 8},
		{4, 3, 5, 8, 7, 3, 3, 5, 4, 1, 5, 4, 6, 9, 8, 4, 4, 6, 5, 2, 6, 5, 7, 1, 9, 5, 5, 7, 6, 3, 7, 6, 8, 2, 1, 6, 6, 8, 7, 4, 8, 7, 9, 3, 2, 7, 7, 9, 8, 5},
		{5, 8, 2, 6, 2, 5, 3, 7, 8, 2, 6, 9, 3, 7, 3, 6, 4, 8, 9, 3, 7, 1, 4, 8, 4, 7, 5, 9, 1, 4, 8, 2, 5, 9, 5, 8, 6, 1, 2, 5, 9, 3, 6, 1, 6, 9, 7, 2, 3, 6},
		{9, 6, 8, 5, 6, 3, 9, 3, 3, 3, 1, 7, 9, 6, 7, 4, 1, 4, 4, 4, 2, 8, 1, 7, 8, 5, 2, 5, 5, 5, 3, 9, 2, 8, 9, 6, 3, 6, 6, 6, 4, 1, 3, 9, 1, 7, 4, 7, 7, 7},
		{3, 5, 3, 2, 3, 4, 1, 3, 5, 9, 4, 6, 4, 3, 4, 5, 2, 4, 6, 1, 5, 7, 5, 4, 5, 6, 3, 5, 7, 2, 6, 8, 6, 5, 6, 7, 4, 6, 8, 3, 7, 9, 7, 6, 7, 8, 5, 7, 9, 4},
		{3, 5, 7, 2, 2, 3, 4, 6, 4, 3, 4, 6, 8, 3, 3, 4, 5, 7, 5, 4, 5, 7, 9, 4, 4, 5, 6, 8, 6, 5, 6, 8, 1, 5, 5, 6, 7, 9, 7, 6, 7, 9, 2, 6, 6, 7, 8, 1, 8, 7},
		{5, 3, 4, 7, 6, 4, 3, 8, 5, 2, 6, 4, 5, 8, 7, 5, 4, 9, 6, 3, 7, 5, 6, 9, 8, 6, 5, 1, 7, 4, 8, 6, 7, 1, 9, 7, 6, 2, 8, 5, 9, 7, 8, 2, 1, 8, 7, 3, 9, 6},
		{3, 4, 2, 5, 3, 5, 1, 7, 4, 3, 4, 5, 3, 6, 4, 6, 2, 8, 5, 4, 5, 6, 4, 7, 5, 7, 3, 9, 6, 5, 6, 7, 5, 8, 6, 8, 4, 1, 7, 6, 7, 8, 6, 9, 7, 9, 5, 2, 8, 7},
		{4, 5, 3, 3, 2, 6, 6, 7, 1, 3, 5, 6, 4, 4, 3, 7, 7, 8, 2, 4, 6, 7, 5, 5, 4, 8, 8, 9, 3, 5, 7, 8, 6, 6, 5, 9, 9, 1, 4, 6, 8, 9, 7, 7, 6, 1, 1, 2, 5, 7},
		{4, 4, 9, 6, 1, 8, 4, 1, 7, 5, 5, 5, 1, 7, 2, 9, 5, 2, 8, 6, 6, 6, 2, 8, 3, 1, 6, 3, 9, 7, 7, 7, 3, 9, 4, 2, 7, 4, 1, 8, 8, 8, 4, 1, 5, 3, 8, 5, 2, 9},
		{4, 6, 2, 4, 6, 1, 6, 9, 1, 5, 5, 7, 3, 5, 7, 2, 7, 1, 2, 6, 6, 8, 4, 6, 8, 3, 8, 2, 3, 7, 7, 9, 5, 7, 9, 4, 9, 3, 4, 8, 8, 1, 6, 8, 1, 5, 1, 4, 5, 9},
		{5, 4, 6, 9, 8, 4, 4, 6, 5, 2, 6, 5, 7, 1, 9, 5, 5, 7, 6, 3, 7, 6, 8, 2, 1, 6, 6, 8, 7, 4, 8, 7, 9, 3, 2, 7, 7, 9, 8, 5, 9, 8, 1, 4, 3, 8, 8, 1, 9, 6},
		{6, 9, 3, 7, 3, 6, 4, 8, 9, 3, 7, 1, 4, 8, 4, 7, 5, 9, 1, 4, 8, 2, 5, 9, 5, 8, 6, 1, 2, 5, 9, 3, 6, 1, 6, 9, 7, 2, 3, 6, 1, 4, 7, 2, 7, 1, 8, 3, 4, 7},
		{1, 7, 9, 6, 7, 4, 1, 4, 4, 4, 2, 8, 1, 7, 8, 5, 2, 5, 5, 5, 3, 9, 2, 8, 9, 6, 3, 6, 6, 6, 4, 1, 3, 9, 1, 7, 4, 7, 7, 7, 5, 2, 4, 1, 2, 8, 5, 8, 8, 8},
		{4, 6, 4, 3, 4, 5, 2, 4, 6, 1, 5, 7, 5, 4, 5, 6, 3, 5, 7, 2, 6, 8, 6, 5, 6, 7, 4, 6, 8, 3, 7, 9, 7, 6, 7, 8, 5, 7, 9, 4, 8, 1, 8, 7, 8, 9, 6, 8, 1, 5},
		{4, 6, 8, 3, 3, 4, 5, 7, 5, 4, 5, 7, 9, 4, 4, 5, 6, 8, 6, 5, 6, 8, 1, 5, 5, 6, 7, 9, 7, 6, 7, 9, 2, 6, 6, 7, 8, 1, 8, 7, 8, 1, 3, 7, 7, 8, 9, 2, 9, 8},
		{6, 4, 5, 8, 7, 5, 4, 9, 6, 3, 7, 5, 6, 9, 8, 6, 5, 1, 7, 4, 8, 6, 7, 1, 9, 7, 6, 2, 8, 5, 9, 7, 8, 2, 1, 8, 7, 3, 9, 6, 1, 8, 9, 3, 2, 9, 8, 4, 1, 7},
		{4, 5, 3, 6, 4, 6, 2, 8, 5, 4, 5, 6, 4, 7, 5, 7, 3, 9, 6, 5, 6, 7, 5, 8, 6, 8, 4, 1, 7, 6, 7, 8, 6, 9, 7, 9, 5, 2, 8, 7, 8, 9, 7, 1, 8, 1, 6, 3, 9, 8},
		{5, 6, 4, 4, 3, 7, 7, 8, 2, 4, 6, 7, 5, 5, 4, 8, 8, 9, 3, 5, 7, 8, 6, 6, 5, 9, 9, 1, 4, 6, 8, 9, 7, 7, 6, 1, 1, 2, 5, 7, 9, 1, 8, 8, 7, 2, 2, 3, 6, 8},
		{5, 5, 1, 7, 2, 9, 5, 2, 8, 6, 6, 6, 2, 8, 3, 1, 6, 3, 9, 7, 7, 7, 3, 9, 4, 2, 7, 4, 1, 8, 8, 8, 4, 1, 5, 3, 8, 5, 2, 9, 9, 9, 5, 2, 6, 4, 9, 6, 3, 1},
		{5, 7, 3, 5, 7, 2, 7, 1, 2, 6, 6, 8, 4, 6, 8, 3, 8, 2, 3, 7, 7, 9, 5, 7, 9, 4, 9, 3, 4, 8, 8, 1, 6, 8, 1, 5, 1, 4, 5, 9, 9, 2, 7, 9, 2, 6, 2, 5, 6, 1},
		{6, 5, 7, 1, 9, 5, 5, 7, 6, 3, 7, 6, 8, 2, 1, 6, 6, 8, 7, 4, 8, 7, 9, 3, 2, 7, 7, 9, 8, 5, 9, 8, 1, 4, 3, 8, 8, 1, 9, 6, 1, 9, 2, 5, 4, 9, 9, 2, 1, 7},
		{7, 1, 4, 8, 4, 7, 5, 9, 1, 4, 8, 2, 5, 9, 5, 8, 6, 1, 2, 5, 9, 3, 6, 1, 6, 9, 7, 2, 3, 6, 1, 4, 7, 2, 7, 1, 8, 3, 4, 7, 2, 5, 8, 3, 8, 2, 9, 4, 5, 8},
		{2, 8, 1, 7, 8, 5, 2, 5, 5, 5, 3, 9, 2, 8, 9, 6, 3, 6, 6, 6, 4, 1, 3, 9, 1, 7, 4, 7, 7, 7, 5, 2, 4, 1, 2, 8, 5, 8, 8, 8, 6, 3, 5, 2, 3, 9, 6, 9, 9, 9},
		{5, 7, 5, 4, 5, 6, 3, 5, 7, 2, 6, 8, 6, 5, 6, 7, 4, 6, 8, 3, 7, 9, 7, 6, 7, 8, 5, 7, 9, 4, 8, 1, 8, 7, 8, 9, 6, 8, 1, 5, 9, 2, 9, 8, 9, 1, 7, 9, 2, 6},
		{5, 7, 9, 4, 4, 5, 6, 8, 6, 5, 6, 8, 1, 5, 5, 6, 7, 9, 7, 6, 7, 9, 2, 6, 6, 7, 8, 1, 8, 7, 8, 1, 3, 7, 7, 8, 9, 2, 9, 8, 9, 2, 4, 8, 8, 9, 1, 3, 1, 9},
		{7, 5, 6, 9, 8, 6, 5, 1, 7, 4, 8, 6, 7, 1, 9, 7, 6, 2, 8, 5, 9, 7, 8, 2, 1, 8, 7, 3, 9, 6, 1, 8, 9, 3, 2, 9, 8, 4, 1, 7, 2, 9, 1, 4, 3, 1, 9, 5, 2, 8},
		{5, 6, 4, 7, 5, 7, 3, 9, 6, 5, 6, 7, 5, 8, 6, 8, 4, 1, 7, 6, 7, 8, 6, 9, 7, 9, 5, 2, 8, 7, 8, 9, 7, 1, 8, 1, 6, 3, 9, 8, 9, 1, 8, 2, 9, 2, 7, 4, 1, 9},
		{6, 7, 5, 5, 4, 8, 8, 9, 3, 5, 7, 8, 6, 6, 5, 9, 9, 1, 4, 6, 8, 9, 7, 7, 6, 1, 1, 2, 5, 7, 9, 1, 8, 8, 7, 2, 2, 3, 6, 8, 1, 2, 9, 9, 8, 3, 3, 4, 7, 9},
	})

	assert.Equal(t, expected, generateWholeMap(parseInput(input), 5))
}

func TestFindLowestRiskLevel(t *testing.T) {
	grid := parseInput(input)

	assert.Equal(t, 40, findLowestRiskLevel(grid))
	assert.Equal(t, 315, findLowestRiskLevel(generateWholeMap(grid, 5)))
}
