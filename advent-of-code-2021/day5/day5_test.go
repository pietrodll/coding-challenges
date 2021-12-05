package day5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

var segments = parseInput(input)

func TestParseInput(t *testing.T) {
	expected := []Segment{
		{Point{0, 9}, Point{5, 9}},
		{Point{8, 0}, Point{0, 8}},
		{Point{9, 4}, Point{3, 4}},
		{Point{2, 2}, Point{2, 1}},
		{Point{7, 0}, Point{7, 4}},
		{Point{6, 4}, Point{2, 0}},
		{Point{0, 9}, Point{2, 9}},
		{Point{3, 4}, Point{1, 4}},
		{Point{0, 0}, Point{8, 8}},
		{Point{5, 5}, Point{8, 2}},
	}

	assert.Equal(t, expected, parseInput(input))
}

func TestInitGrid(t *testing.T) {
	grid := initGrid(segments)

	assert.Equal(t, 10, grid.height)
	assert.Equal(t, 10, grid.width)
}

func TestCountOverlappingVerticalAndHorizontal(t *testing.T) {
	assert.Equal(t, 5, countOverlappingVerticalAndHorizontal(segments))
}

func TestCountOverlapping(t *testing.T) {
	assert.Equal(t, 12, countOverlapping(segments))
}
