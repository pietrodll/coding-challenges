package day17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	assert.Equal(t, TargetArea{20, 30, -10, -5}, parseInput("target area: x=20..30, y=-10..-5"))
}

func TestIsValidInitialVelocity(t *testing.T) {
	target := TargetArea{20, 30, -10, -5}

	assert.True(t, isValidInitialVelocity(target, 7, 2))
	assert.True(t, isValidInitialVelocity(target, 6, 3))
	assert.True(t, isValidInitialVelocity(target, 9, 0))
	assert.False(t, isValidInitialVelocity(target, 17, -4))
}

func TestFindVxFromFinalPosition(t *testing.T) {
	vx, ok := findVxFromFinalPosition(10)
	assert.True(t, ok)
	assert.Equal(t, 4, vx)

	_, ok = findVxFromFinalPosition(11)
	assert.False(t, ok)
}

func TestFindHighestPossibleAltitute(t *testing.T) {
	target := TargetArea{20, 30, -10, -5}
	assert.Equal(t, 45, findHighestPossibleAltitute(target))
}

func TestCountValidInitialVelocities(t *testing.T) {
	target := TargetArea{20, 30, -10, -5}
	assert.Equal(t, 112, countValidInitialVelocities(target))
}
