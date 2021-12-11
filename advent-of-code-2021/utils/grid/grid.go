package grid

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/pietrodll/aoc2021/utils/base"
)

type Grid struct {
	values [][]int
	Height int
	Width  int
}

// Constructor
func NewGrid(values [][]int) Grid {
	height := len(values)
	width := len(values[0])

	valuesCopy := make([][]int, height)

	for i, line := range values {
		if len(line) != width {
			panic(fmt.Errorf("values are not a grid: length of line %d is %d", i, len(line)))
		}

		lineCopy := make([]int, width)
		copy(lineCopy, line)
		valuesCopy[i] = lineCopy
	}

	return Grid{valuesCopy, height, width}
}

func NewGridFromString(data string, rowSep string, valSep string) Grid {
	lines := strings.Split(data, rowSep)
	height := len(lines)
	width := 0

	grid := make([][]int, height)

	for i, line := range lines {
		strValues := strings.Split(line, valSep)

		if i == 0 {
			width = len(strValues)
		} else if len(strValues) != width {
			panic(fmt.Errorf("values are not a grid: length of line %d is %d", i, len(strValues)))
		}

		row := make([]int, width)
		for j, strVal := range strValues {
			val, err := strconv.Atoi(strVal)

			if err != nil {
				panic(err)
			}

			row[j] = val
		}

		grid[i] = row
	}

	return Grid{grid, height, width}
}

// Implement coder interface
func (g *Grid) Encode(point base.Codable) int {
	switch point := point.(type) {
	case GridPoint:
		return point.I*g.Width + point.J
	default:
		panic(errors.New("a Grid can only encode GridPoint"))
	}
}

func (g *Grid) Decode(encoded int) base.Codable {
	return GridPoint{encoded / g.Width, encoded % g.Width}
}

// Adjacent points
func (g *Grid) FindAdjacentPoints(point GridPoint) []GridPoint {
	i, j := point.I, point.J
	points := make([]GridPoint, 0)

	if i > 0 {
		points = append(points, GridPoint{i - 1, j})
	}

	if i < g.Height-1 {
		points = append(points, GridPoint{i + 1, j})
	}

	if j > 0 {
		points = append(points, GridPoint{i, j - 1})
	}

	if j < g.Width-1 {
		points = append(points, GridPoint{i, j + 1})
	}

	return points
}

func (g *Grid) FindAdjacentPointsWithDiagonals(point GridPoint) []GridPoint {
	i, j := point.I, point.J
	points := make([]GridPoint, 0, 4)

	if i > 0 {
		points = append(points, GridPoint{i - 1, j})

		if j > 0 {
			points = append(points, GridPoint{i - 1, j - 1})
		}

		if j < g.Width-1 {
			points = append(points, GridPoint{i - 1, j + 1})
		}
	}

	if i < g.Height-1 {
		points = append(points, GridPoint{i + 1, j})

		if j > 0 {
			points = append(points, GridPoint{i + 1, j - 1})
		}

		if j < g.Width-1 {
			points = append(points, GridPoint{i + 1, j + 1})
		}

	}

	if j > 0 {
		points = append(points, GridPoint{i, j - 1})
	}

	if j < g.Width-1 {
		points = append(points, GridPoint{i, j + 1})
	}

	return points
}

func (g *Grid) GetValue(point GridPoint) int {
	return g.values[point.I][point.J]
}

func (g *Grid) GetPtr(point GridPoint) *int {
	return &g.values[point.I][point.J]
}

func (g *Grid) SetValue(point GridPoint, value int) {
	g.values[point.I][point.J] = value
}

func (g *Grid) StreamPoints() chan GridPoint {
	stream := make(chan GridPoint)

	go func() {
		for i, line := range g.values {
			for j := range line {
				stream <- GridPoint{i, j}
			}
		}

		close(stream)
	}()

	return stream
}

func (g *Grid) Copy() Grid {
	return NewGrid(g.values)
}
