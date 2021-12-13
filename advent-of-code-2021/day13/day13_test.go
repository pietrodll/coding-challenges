package day13

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`

func TestParseInput(t *testing.T) {
	expectedPoints := []Point{
		{6, 10},
		{0, 14},
		{9, 10},
		{0, 3},
		{10, 4},
		{4, 11},
		{6, 0},
		{6, 12},
		{4, 1},
		{0, 13},
		{10, 12},
		{3, 4},
		{3, 0},
		{8, 4},
		{1, 10},
		{2, 14},
		{8, 10},
		{9, 0},
	}
	expectedFolds := []Fold{YFold{7}, XFold{5}}

	points, folds := parseInput(input)
	assert.Equal(t, expectedPoints, points)
	assert.Equal(t, expectedFolds, folds)
}

func TestFolds(t *testing.T) {
	assert.Equal(t, Point{0, 0}, XFold{2}.transform(Point{4, 0}))
	assert.Equal(t, Point{1, 0}, XFold{2}.transform(Point{1, 0}))
	assert.Equal(t, Point{1, 2}, YFold{3}.transform(Point{1, 4}))
	assert.Equal(t, Point{1, 1}, YFold{3}.transform(Point{1, 1}))
}

func TestExecuteFolds(t *testing.T) {
	points := []Point{
		{6, 10},
		{0, 14},
		{9, 10},
		{0, 3},
		{10, 4},
		{4, 11},
		{6, 0},
		{6, 12},
		{4, 1},
		{0, 13},
		{10, 12},
		{3, 4},
		{3, 0},
		{8, 4},
		{1, 10},
		{2, 14},
		{8, 10},
		{9, 0},
	}
	folds := []Fold{YFold{7}, XFold{5}}
	expected := []Point{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 0},
		{0, 1},
		{0, 2},
		{0, 3},
		{4, 1},
		{4, 2},
		{4, 3},
		{0, 4},
		{1, 4},
		{2, 4},
		{3, 4},
		{4, 4},
	}
	assert.ElementsMatch(t, expected, executeFolds(points, folds))
}

func TestCountVisibleAfterFirstFold(t *testing.T) {
	assert.Equal(t, 17, countPointsAfterFirstFold(parseInput(input)))
}

func TestDisplay(t *testing.T) {
	points := []Point{{0, 0}, {1, 0}, {2, 0}, {2, 2}}
	assert.Equal(t, "###\n...\n..#", displayPoints(points))
}
