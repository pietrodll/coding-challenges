package day9

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/pietrodll/aoc2021/utils/collections"
)

type Grid struct {
	grid   [][]int
	Height int
	Width  int
}

type GridPoint struct {
	I int
	J int
}

func (g *Grid) decode(hash int) GridPoint {
	i := hash / g.Width
	j := hash % g.Width

	return GridPoint{i, j}
}

func (g *Grid) encode(point GridPoint) int {
	return point.I*g.Width + point.J
}

func (g *Grid) getVal(point GridPoint) int {
	return g.grid[point.I][point.J]
}

func parseInput(input string) Grid {
	lines := strings.Split(input, "\n")
	grid := make([][]int, len(lines))

	for i, line := range lines {
		gridLine := make([]int, len(line))

		for j, strVal := range strings.Split(line, "") {
			val, err := strconv.Atoi(strVal)

			if err != nil {
				panic(err)
			}

			gridLine[j] = val
		}

		grid[i] = gridLine
	}

	return Grid{grid, len(grid), len(grid[0])}
}

func (g *Grid) findAdjacentPoints(point GridPoint) []GridPoint {
	i, j := point.I, point.J
	points := make([]GridPoint, 0, 4)

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

func (g *Grid) findLowPoints() []GridPoint {
	lowPoints := make([]GridPoint, 0)

	for i, line := range g.grid {
		for j, val := range line {
			adjacent := g.findAdjacentPoints(GridPoint{i, j})
			isLow := true

			for _, point := range adjacent {
				isLow = isLow && (val < g.getVal(point))
			}

			if isLow {
				lowPoints = append(lowPoints, GridPoint{i, j})
			}
		}
	}

	return lowPoints
}

func (g *Grid) totalRiskLevel() int {
	riskLevel := 0

	for _, point := range g.findLowPoints() {
		riskLevel += 1 + g.grid[point.I][point.J]
	}

	return riskLevel
}

func (g *Grid) findBasins() [][]GridPoint {
	visited := make([]bool, g.Height*g.Width)

	lowPoints := g.findLowPoints()
	basins := make([][]GridPoint, len(lowPoints))

	for i, lowPoint := range lowPoints {
		encodedLowPoint := g.encode(lowPoint)

		if !visited[encodedLowPoint] {
			// breadth-first search starting from the low point to explore the basin
			basin := make([]GridPoint, 1)
			basin[0] = lowPoint

			to_visit := collections.NewIntQueue(encodedLowPoint)
			visited[encodedLowPoint] = true

			for !to_visit.IsEmpty() {
				point := g.decode(to_visit.Dequeue())
				pointVal := g.getVal(point)

				for _, neighbor := range g.findAdjacentPoints(point) {
					encodedNeighbor := g.encode(neighbor)
					val := g.getVal(neighbor)

					if val >= pointVal && val < 9 && !visited[encodedNeighbor] {
						to_visit.Enqueue(encodedNeighbor)
						basin = append(basin, neighbor)
						visited[encodedNeighbor] = true
					}

				}
			}

			basins[i] = basin
		}
	}

	return basins
}

func Run(input string) {
	grid := parseInput(input)

	fmt.Println("Total risk level:", grid.totalRiskLevel())

	basins := grid.findBasins()
	basinSizes := make([]int, len(basins))

	for i, basin := range basins {
		basinSizes[i] = len(basin)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))

	fmt.Println("Three largest basins multiplied:", basinSizes[0]*basinSizes[1]*basinSizes[2])
}
