package day11

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pietrodll/aoc2021/utils/collections"
)

func parseInput(input string) [][]int {
	lines := strings.Split(input, "\n")

	grid := make([][]int, len(lines))

	for i, line := range lines {
		gridLine := make([]int, len(line))

		for j, strVal := range line {
			val, err := strconv.Atoi(string(strVal))

			if err != nil {
				panic(err)
			}

			gridLine[j] = val
		}

		grid[i] = gridLine
	}

	return grid
}

type OctopusSimulator struct {
	grid   [][]int
	height int
	width  int
}

func newOctopusSimulator(grid [][]int) OctopusSimulator {
	height, width := len(grid), len(grid[0])

	gridCopy := make([][]int, height)

	for i, line := range grid {
		gridCopy[i] = make([]int, width)
		copy(gridCopy[i], line)
	}

	return OctopusSimulator{gridCopy, height, width}
}

type GridPoint struct {
	I int
	J int
}

func (sim *OctopusSimulator) findAdjacentPointsWithDiagonals(point GridPoint) []GridPoint {
	i, j := point.I, point.J
	points := make([]GridPoint, 0, 4)

	if i > 0 {
		points = append(points, GridPoint{i - 1, j})

		if j > 0 {
			points = append(points, GridPoint{i - 1, j - 1})
		}

		if j < sim.width-1 {
			points = append(points, GridPoint{i - 1, j + 1})
		}
	}

	if i < sim.height-1 {
		points = append(points, GridPoint{i + 1, j})

		if j > 0 {
			points = append(points, GridPoint{i + 1, j - 1})
		}

		if j < sim.width-1 {
			points = append(points, GridPoint{i + 1, j + 1})
		}

	}

	if j > 0 {
		points = append(points, GridPoint{i, j - 1})
	}

	if j < sim.width-1 {
		points = append(points, GridPoint{i, j + 1})
	}

	return points
}

func (sim *OctopusSimulator) encode(point GridPoint) int {
	return point.I*sim.width + point.J
}

func (sim *OctopusSimulator) decode(encoded int) GridPoint {
	i := encoded / sim.width
	j := encoded % sim.width
	return GridPoint{i, j}
}

func (sim *OctopusSimulator) nextStep() int {
	flashedCount := 0

	willFlash := collections.NewIntSet()

	// raise all the energy levels by 1
	for i, line := range sim.grid {
		for j := range line {
			sim.grid[i][j] += 1

			if sim.grid[i][j] > 9 {
				willFlash.Add(sim.encode(GridPoint{i, j}))
			}
		}
	}

	// flash
	hasFlashed := collections.NewIntSet()
	for !willFlash.IsEmpty() {
		encoded := willFlash.Pop()
		point := sim.decode(encoded)
		flashedCount += 1
		sim.grid[point.I][point.J] = 0
		hasFlashed.Add(encoded)

		for _, adjacent := range sim.findAdjacentPointsWithDiagonals(point) {
			encodedAdjacent := sim.encode(adjacent)

			if !hasFlashed.Contains(encodedAdjacent) {
				sim.grid[adjacent.I][adjacent.J] += 1

				if sim.grid[adjacent.I][adjacent.J] > 9 {
					willFlash.Add(encodedAdjacent)
				}
			}

		}
	}

	return flashedCount
}

func (sim *OctopusSimulator) totalFlashes(steps int) int {
	totFlashes := 0

	for s := 0; s < steps; s++ {
		totFlashes += sim.nextStep()
	}

	return totFlashes
}

func (sim *OctopusSimulator) computeUntilSynchronized() int {
	step := 0
	gridSize := sim.height * sim.width

	for flashCount := 0; flashCount != gridSize; flashCount = sim.nextStep() {
		step++
	}

	return step
}

func Run(input string) {
	grid := parseInput(input)
	sim := newOctopusSimulator(grid)

	fmt.Println("Total flashes after 100 steps:", sim.totalFlashes(100))

	sim = newOctopusSimulator(grid)
	fmt.Println("Steps necessary to synchronize:", sim.computeUntilSynchronized())
}
