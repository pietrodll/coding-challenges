package day7

import (
	"fmt"
	"math"
	"sort"

	"github.com/pietrodll/aoc2021/utils/parse"
)

func parseInput(input string) []int {
	return parse.ParseIntegers(input, ",")
}

func findMinimumManhattanDistancePoint(positions []int) int {
	positionsCopy := make([]int, len(positions))
	copy(positionsCopy, positions)
	sort.Ints(positionsCopy)

	return positionsCopy[len(positionsCopy)/2]
}

func computeConstantFuelComsumption(positions []int, target int) int {
	consumption := 0

	for _, pos := range positions {
		if pos > target {
			consumption += pos - target
		} else {
			consumption += target - pos
		}
	}

	return consumption
}

func computeMinimumConstantFuelConsumption(positions []int) int {
	return computeConstantFuelComsumption(positions, findMinimumManhattanDistancePoint(positions))
}

func computeIncreasingFuelConsumption(positions []int, target int) int {
	consumption := 0

	for _, pos := range positions {
		d := pos - target

		if d < 0 {
			d = -d
		}

		consumption += (d * (d + 1)) / 2
	}

	return consumption
}

func findMinimumDistancePoint(positions []int) int {
	avg := float64(0)

	for _, pos := range positions {
		avg += float64(pos)
	}

	avg /= float64(len(positions))

	return int(math.Round(avg))
}

// Computes the minimum consumption for the second part of the challenge.
// If a crab moves from point A to point B, its consumption is:
// C(A, B) = 1 + 2 + ... + |A - B| = ((A - B)^2 + |A - B|) / 2
//
// Note that |A - B| <= C(A, B) <= (A - B)^2.
// Therefore, the optimal target point lies between the minimum Manhattan distance point and the barycenter (optimal for L2 distance)
func computeMinimumIncreasingFuelConsumption(positions []int) int {
	start := findMinimumManhattanDistancePoint(positions)
	end := findMinimumDistancePoint(positions)

	if start > end {
		start, end = end, start
	}

	minConsumption := computeIncreasingFuelConsumption(positions, end)

	for target := start; target < end; target++ {
		consumption := computeIncreasingFuelConsumption(positions, target)

		if consumption < minConsumption {
			minConsumption = consumption
		}
	}

	return minConsumption
}

func Run(input string) {
	positions := parseInput(input)

	fmt.Println("Minimum fuel consumption:", computeMinimumConstantFuelConsumption(positions))
	fmt.Println("Minimum fuel consumption (increasing):", computeMinimumIncreasingFuelConsumption(positions))
}
