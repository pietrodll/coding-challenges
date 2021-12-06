package day6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var initialState = []int{3, 4, 3, 1, 2}

func TestInitStateTracker(t *testing.T) {
	tracker := initStateTracker(initialState, 8, 6)
	assert.Equal(t, []int{0, 1, 1, 2, 1, 0, 0, 0, 0}, tracker.countByTimer)
}

func TestNextState(t *testing.T) {
	tracker := initStateTracker(initialState, 8, 6)

	tracker.nextState()
	assert.Equal(t, []int{1, 1, 2, 1, 0, 0, 0, 0, 0}, tracker.countByTimer)

	tracker.nextState()
	assert.Equal(t, []int{1, 2, 1, 0, 0, 0, 1, 0, 1}, tracker.countByTimer)
}
