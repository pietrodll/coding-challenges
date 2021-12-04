package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `forward 5
down 5
forward 8
up 3
down 8
forward 2`

func TestDay2Part1(t *testing.T) {
	instructions := parseInstructions(input)

	position := Position{0, 0}

	executeInstructions(instructions, &position)

	assert.Equal(t, 15, position.Horiz)
	assert.Equal(t, 10, position.Depth)
}

func TestDay2Part2(t *testing.T) {
	instructions := parseInstructions(input)

	position := PositionWithAim{Position{0, 0}, 0}

	executeInstructions(instructions, &position)

	assert.Equal(t, 15, position.Horiz)
	assert.Equal(t, 60, position.Depth)
	assert.Equal(t, 10, position.Aim)
}
