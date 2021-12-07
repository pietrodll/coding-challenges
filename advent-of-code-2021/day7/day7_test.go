package day7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var positions = []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

func TestFindMinimumDistancePoint(t *testing.T) {
	assert.Equal(t, 2, findMinimumManhattanDistancePoint(positions))
}

func TestComputeConstantFuelConsumption(t *testing.T) {
	assert.Equal(t, 37, computeConstantFuelComsumption(positions, 2))
}

func TestComputeIncreasingFuelConsumption(t *testing.T) {
	assert.Equal(t, 206, computeIncreasingFuelConsumption(positions, 2))
	assert.Equal(t, 168, computeIncreasingFuelConsumption(positions, 5))
}

func TestComputeMinimumIncreasingFuelConsumption(t *testing.T) {
	assert.Equal(t, 168, computeMinimumIncreasingFuelConsumption(positions))
}
