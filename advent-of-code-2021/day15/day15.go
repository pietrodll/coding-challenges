package day15

import (
	"fmt"

	"github.com/pietrodll/aoc2021/utils/collections"
	"github.com/pietrodll/aoc2021/utils/grid"
)

func parseInput(input string) grid.Grid {
	return grid.NewGridFromString(input, "\n", "")
}

func generateWholeMap(g grid.Grid, scale int) grid.Grid {
	values := make([][]int, scale*g.Height)

	for i := range values {
		line := make([]int, scale*g.Width)

		for j := range line {
			baseGridPoint := grid.GridPoint{I: i % g.Height, J: j % g.Width}
			value := g.GetValue(baseGridPoint) + (i / g.Height) + (j / g.Width)
			value = ((value - 1) % 9) + 1

			line[j] = value
		}

		values[i] = line
	}

	return grid.NewGrid(values)
}

const Infinity = int((^uint(0)) >> 1)

type PointWithDistance struct {
	grid.GridPoint
	distances *map[grid.GridPoint]int
}

func (p PointWithDistance) GetPriority() int {
	return (*p.distances)[p.GridPoint]
}

func (p PointWithDistance) Key() interface{} {
	return p.GridPoint
}

func findLowestRiskLevel(g grid.Grid) int {
	distance := make(map[grid.GridPoint]int)
	visited := make(map[grid.GridPoint]bool)
	q := collections.NewPriorityQueue()

	for point := range g.StreamPoints() {
		distance[point] = Infinity
		q.Add(PointWithDistance{point, &distance})
	}

	distance[grid.GridPoint{I: 0, J: 0}] = 0

	for !q.IsEmpty() {
		u := q.Pop().(PointWithDistance)
		visited[u.GridPoint] = true

		for _, v := range g.FindAdjacentPoints(u.GridPoint) {
			if !visited[v] {
				if alt := distance[u.GridPoint] + g.GetValue(v); alt < distance[v] {
					distance[v] = alt
					q.DecreasedPriority(PointWithDistance{v, &distance})
				}
			}
		}
	}

	return distance[grid.GridPoint{I: g.Height - 1, J: g.Width - 1}]
}

func Run(input string) {
	grid := parseInput(input)
	fmt.Println("Minimum total risk level:", findLowestRiskLevel(grid))

	wholeMap := generateWholeMap(grid, 5)
	fmt.Println("Minimum total risk level for whole map:", findLowestRiskLevel(wholeMap))
}
