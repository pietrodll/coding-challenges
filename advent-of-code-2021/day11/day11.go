package day11

import (
	"fmt"

	"github.com/pietrodll/aoc2021/utils/collections"
	"github.com/pietrodll/aoc2021/utils/grid"
)

func parseInput(input string) grid.Grid {
	return grid.NewGridFromString(input, "\n", "")
}

type OctopusSimulator struct {
	g grid.Grid
}

func newOctopusSimulator(g grid.Grid) OctopusSimulator {
	return OctopusSimulator{g.Copy()}
}

func (sim *OctopusSimulator) nextStep() int {
	flashedCount := 0

	willFlash := collections.NewCodableSet(&sim.g)

	// raise all the energy levels by 1
	for point := range sim.g.StreamPoints() {
		*sim.g.GetPtr(point)++

		if sim.g.GetValue(point) > 9 {
			willFlash.Add(point)
		}
	}

	// flash
	hasFlashed := collections.NewCodableSet(&sim.g)
	for !willFlash.IsEmpty() {
		point := willFlash.Pop().(grid.GridPoint)
		flashedCount += 1
		sim.g.SetValue(point, 0)
		hasFlashed.Add(point)

		for _, adjacent := range sim.g.FindAdjacentPointsWithDiagonals(point) {
			if !hasFlashed.Contains(adjacent) {
				*sim.g.GetPtr(adjacent)++

				if sim.g.GetValue(adjacent) > 9 {
					willFlash.Add(adjacent)
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
	gridSize := sim.g.Height * sim.g.Width

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
