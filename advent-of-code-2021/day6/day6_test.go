package day6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	assert.Equal(t, []int{3, 4, 3, 1, 2}, parseInput("3,4,3,1,2"))
}

var initialState = []int{3, 4, 3, 1, 2}

func TestInitStateTracker(t *testing.T) {
	tracker := initStateTracker(initialState, 8, 6)
	assert.Equal(t, []int{0, 1, 1, 2, 1, 0, 0, 0, 0}, tracker.countByTimer)

	assert.PanicsWithError(t, "lanternfish timer cannot be 4", func() {
		initStateTracker(initialState, 3, 2)
	})
}

func TestNextState(t *testing.T) {
	tracker := initStateTracker(initialState, 8, 6)

	tracker.nextState()
	assert.Equal(t, []int{1, 1, 2, 1, 0, 0, 0, 0, 0}, tracker.countByTimer)

	tracker.nextState()
	assert.Equal(t, []int{1, 2, 1, 0, 0, 0, 1, 0, 1}, tracker.countByTimer)
}

func TestSimulateGrowth(t *testing.T) {
	tracker := initStateTracker(initialState, 8, 6)

	tracker.simulateGrowth(2)
	assert.Equal(t, []int{1, 2, 1, 0, 0, 0, 1, 0, 1}, tracker.countByTimer)
}

func TestCountFishes(t *testing.T) {
	tracker := initStateTracker(initialState, 8, 6)
	assert.Equal(t, 5, tracker.countFishes())
}
