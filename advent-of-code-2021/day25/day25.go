package day25

import (
	"fmt"
	"strings"

	"github.com/pietrodll/aoc2021/utils/grid"
)

type CucumberGrid struct {
	height, width int
	// 1 = >, 2 = v
	cucumbers map[grid.GridPoint]uint8
}

func parseInput(input string) CucumberGrid {
	lines := strings.Split(input, "\n")
	cucumbers := make(map[grid.GridPoint]uint8)
	height := len(lines)
	width := len(lines[0])

	for i, line := range lines {
		for j, char := range line {
			if char == '>' {
				cucumbers[grid.GridPoint{I: i, J: j}] = uint8(1)
			} else if char == 'v' {
				cucumbers[grid.GridPoint{I: i, J: j}] = uint8(2)
			}
		}
	}

	return CucumberGrid{height: height, width: width, cucumbers: cucumbers}
}

func nextStep(g CucumberGrid) (CucumberGrid, bool) {
	hasMoved := false
	newCucumbers := make(map[grid.GridPoint]uint8)
	tempCucumbers := make(map[grid.GridPoint]uint8)

	for pos, val := range g.cucumbers {
		if val == uint8(1) {
			// if > try to move it
			nextPos := grid.GridPoint{I: pos.I, J: (pos.J + 1) % g.width}

			if _, present := g.cucumbers[nextPos]; !present {
				newCucumbers[nextPos] = val
				tempCucumbers[nextPos] = val
				hasMoved = true
			} else {
				newCucumbers[pos] = val
				tempCucumbers[pos] = val
			}
		} else {
			// otherwise, just copy to temp
			tempCucumbers[pos] = val
		}
	}

	for pos, val := range g.cucumbers {
		if val == uint8(2) {
			// if v try to move it
			nextPos := grid.GridPoint{I: (pos.I + 1) % g.height, J: pos.J}

			if _, present := tempCucumbers[nextPos]; !present {
				newCucumbers[nextPos] = val
				hasMoved = true
			} else {
				newCucumbers[pos] = val
			}
		}
	}

	return CucumberGrid{g.height, g.width, newCucumbers}, hasMoved
}

func findFirstStoppingStep(g CucumberGrid) int {
	step := 1
	nextGrid, hasMoved := nextStep(g)

	for hasMoved {
		nextGrid, hasMoved = nextStep(nextGrid)
		step++
	}

	return step
}

func Run(input string) {
	cucumberGrid := parseInput(input)

	fmt.Println("First stopping step:", findFirstStoppingStep(cucumberGrid))
}
