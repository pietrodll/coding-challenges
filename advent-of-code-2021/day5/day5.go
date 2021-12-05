package day5

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Segment struct {
	Start Point
	End   Point
}

func (s *Segment) getOrderedPoints() (Point, Point) {
	if s.Start.X < s.End.X || (s.Start.X == s.End.X && s.Start.Y <= s.End.Y) {
		return s.Start, s.End
	}

	return s.End, s.Start
}

type Grid struct {
	grid   [][]int
	height int
	width  int
}

func parseInput(input string) []Segment {
	segmentDescriptions := strings.Split(input, "\n")
	regex := regexp.MustCompile(`\d+`)

	segments := make([]Segment, len(segmentDescriptions))

	for i, segmentDesc := range segmentDescriptions {
		numberStr := regex.FindAllString(segmentDesc, 4)

		if numberStr == nil || len(numberStr) != 4 {
			panic(fmt.Errorf("cannot parse %s", segmentDesc))
		}

		numbers := make([]int, 4)

		for j, strVal := range numberStr {
			val, err := strconv.Atoi(strVal)

			if err != nil {
				panic(err)
			}

			numbers[j] = val
		}

		segments[i] = Segment{
			Point{numbers[0], numbers[1]},
			Point{numbers[2], numbers[3]},
		}
	}

	return segments
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func initGrid(segments []Segment) Grid {
	// find max value in all segments to determine the size of the grid
	maxX, maxY := 0, 0

	for _, segment := range segments {
		segMaxX := max(segment.Start.X, segment.End.X)
		segMaxY := max(segment.Start.Y, segment.End.Y)

		if segMaxX > maxX {
			maxX = segMaxX
		}

		if segMaxY > maxY {
			maxY = segMaxY
		}
	}

	grid := make([][]int, maxX+1)

	for i := range grid {
		grid[i] = make([]int, maxY+1)
	}

	return Grid{grid, maxX + 1, maxY + 1}
}

func (g *Grid) registerSegment(segment Segment) {
	start, end := segment.getOrderedPoints()
	currentX, currentY := start.X, start.Y

	stepX := 1 // we ordered the points, therefore stepX is always 1
	stepY := 1

	if end.Y < start.Y {
		stepY = -1
	}

	for (end.X-currentX)*stepX > 0 || (end.Y-currentY)*stepY > 0 {
		g.grid[currentX][currentY] += 1

		if (end.X-currentX)*stepX > 0 {
			currentX += stepX
		}

		if (end.Y-currentY)*stepY > 0 {
			currentY += stepY
		}
	}

	g.grid[currentX][currentY] += 1
}

func (g *Grid) countOverlapping() int {
	cnt := 0

	for i, row := range g.grid {
		for j := range row {
			if g.grid[i][j] > 1 {
				cnt += 1
			}
		}
	}

	return cnt
}

func countOverlappingVerticalAndHorizontal(segments []Segment) int {
	grid := initGrid(segments)

	for _, segment := range segments {
		if segment.Start.X == segment.End.X || segment.Start.Y == segment.End.Y {
			grid.registerSegment(segment)
		}
	}

	return grid.countOverlapping()
}

func countOverlapping(segments []Segment) int {
	grid := initGrid(segments)

	for _, segment := range segments {
		grid.registerSegment(segment)
	}

	return grid.countOverlapping()
}

func Run(input string) {
	segments := parseInput(input)

	fmt.Println("Overlapping count:", countOverlappingVerticalAndHorizontal(segments))
	fmt.Println("Overlapping count with diagonals:", countOverlapping(segments))
}
