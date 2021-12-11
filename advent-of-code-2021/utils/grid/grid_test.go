package grid

import (
	"testing"

	"github.com/pietrodll/aoc2021/utils/base"
	"github.com/stretchr/testify/assert"
)

func TestGridConstructor(t *testing.T) {
	values := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	grid := NewGrid(values)

	assert.Equal(t, 3, grid.Height)
	assert.Equal(t, 3, grid.Width)
	assert.Equal(t, values, grid.values)

	// check that we create a copy of values
	grid.values[0][0] = 5
	assert.NotEqual(t, values, grid.values)

	badValues := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8},
	}

	assert.PanicsWithError(t, "values are not a grid: length of line 2 is 2", func() {
		NewGrid(badValues)
	})
}

type DummyCodable struct{}

func (d DummyCodable) Encode(coder base.Coder) int {
	return 0
}

func TestGridEncoding(t *testing.T) {
	grid := NewGrid([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	assert.Equal(t, 2, grid.Encode(GridPoint{0, 2}))
	assert.Equal(t, 8, grid.Encode(GridPoint{2, 2}))

	assert.PanicsWithError(t, "a Grid can only encode GridPoint", func() {
		grid.Encode(DummyCodable{})
	})

	assert.Equal(t, GridPoint{1, 1}, grid.Decode(4))
}

func TestAdjacentPoints(t *testing.T) {
	values := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	grid := NewGrid(values)

	assert.ElementsMatch(t, []GridPoint{{0, 1}, {1, 0}}, grid.FindAdjacentPoints(GridPoint{0, 0}))
	assert.ElementsMatch(t, []GridPoint{{0, 1}, {1, 0}, {1, 2}, {2, 1}}, grid.FindAdjacentPoints(GridPoint{1, 1}))
}

func TestAdjacentPointsWithDiagonals(t *testing.T) {
	grid := NewGrid([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	})

	assert.ElementsMatch(t, []GridPoint{{0, 1}, {1, 0}, {1, 1}}, grid.FindAdjacentPointsWithDiagonals(GridPoint{0, 0}))
	assert.ElementsMatch(t, []GridPoint{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 2}, {2, 0}, {2, 1}, {2, 2}}, grid.FindAdjacentPointsWithDiagonals(GridPoint{1, 1}))
}

func TestGettersAndSetters(t *testing.T) {
	grid := NewGrid([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	})

	point := GridPoint{0, 1}
	assert.Equal(t, 2, grid.GetValue(point))
	assert.Equal(t, 2, *grid.GetPtr(point))

	*grid.GetPtr(point)++
	assert.Equal(t, 3, grid.GetValue(point))

	grid.SetValue(point, 5)
	assert.Equal(t, 5, grid.GetValue(point))
}

func TestStreamPoints(t *testing.T) {
	grid := NewGrid([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	})

	counter := 0

	for range grid.StreamPoints() {
		counter++
	}

	assert.Equal(t, 16, counter)
}

func TestCopy(t *testing.T) {
	grid := NewGrid([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	})

	copy := grid.Copy()

	point := GridPoint{0, 0}
	grid.SetValue(point, 42)
	assert.Equal(t, 42, grid.GetValue(point))
	assert.Equal(t, 1, copy.GetValue(point))
}

func TestNewGridFromString(t *testing.T) {
	data := "123-456-789"
	expected := NewGrid([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	assert.Equal(t, expected, NewGridFromString(data, "-", ""))

	data = "123-4x5-789"
	assert.Panics(t, func() {
		NewGridFromString(data, "-", "")
	})

	data = "123-456-78"
	assert.PanicsWithError(t, "values are not a grid: length of line 2 is 2", func() {
		NewGridFromString(data, "-", "")
	})
}
