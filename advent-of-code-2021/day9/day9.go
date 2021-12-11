package day9

import (
	"fmt"
	"sort"

	"github.com/pietrodll/aoc2021/utils/collections"
	"github.com/pietrodll/aoc2021/utils/grid"
)

func parseInput(input string) grid.Grid {
	return grid.NewGridFromString(input, "\n", "")
}

func findLowPoints(g *grid.Grid) []grid.GridPoint {
	lowPoints := make([]grid.GridPoint, 0)

	for point := range g.StreamPoints() {
		adjacent := g.FindAdjacentPoints(point)
		val := g.GetValue(point)
		isLow := true

		for _, adj := range adjacent {
			isLow = isLow && (val < g.GetValue(adj))
		}

		if isLow {
			lowPoints = append(lowPoints, point)
		}
	}

	return lowPoints
}

func totalRiskLevel(g *grid.Grid) int {
	riskLevel := 0

	for _, point := range findLowPoints(g) {
		riskLevel += 1 + g.GetValue(point)
	}

	return riskLevel
}

func findBasins(g *grid.Grid) [][]grid.GridPoint {
	visited := collections.NewCodableSet(g)

	lowPoints := findLowPoints(g)
	basins := make([][]grid.GridPoint, len(lowPoints))

	for i, lowPoint := range lowPoints {
		if !visited.Contains(lowPoint) {
			// breadth-first search starting from the low point to explore the basin
			basin := make([]grid.GridPoint, 1)
			basin[0] = lowPoint

			toVisit := collections.NewQueue(lowPoint)
			visited.Add(lowPoint)

			for !toVisit.IsEmpty() {
				point := toVisit.Dequeue().(grid.GridPoint)
				pointVal := g.GetValue(point)

				for _, neighbor := range g.FindAdjacentPoints(point) {
					val := g.GetValue(neighbor)

					if val >= pointVal && val < 9 && !visited.Contains(neighbor) {
						toVisit.Enqueue(neighbor)
						basin = append(basin, neighbor)
						visited.Add(neighbor)
					}
				}
			}

			basins[i] = basin
		}
	}

	return basins
}

func findAndMultiplyThreeLargestBasins(g *grid.Grid) int {
	basins := findBasins(g)
	basinSizes := make([]int, len(basins))

	for i, basin := range basins {
		basinSizes[i] = len(basin)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	return basinSizes[0] * basinSizes[1] * basinSizes[2]
}

func Run(input string) {
	grid := parseInput(input)

	fmt.Println("Total risk level:", totalRiskLevel(&grid))

	fmt.Println("Three largest basins multiplied:", findAndMultiplyThreeLargestBasins(&grid))
}
