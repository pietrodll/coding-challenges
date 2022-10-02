package day24

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecrementDigits(t *testing.T) {
	digits := []int{9, 9, 9}

	decrementDigits(digits)
	assert.Equal(t, []int{9, 9, 8}, digits)

	digits = []int{9, 9, 1}
	decrementDigits(digits)
	assert.Equal(t, []int{9, 8, 9}, digits)

	digits = []int{9, 1, 1}
	decrementDigits(digits)
	assert.Equal(t, []int{8, 9, 9}, digits)

	digits = []int{1, 1, 1}
	decrementDigits(digits)
	assert.Equal(t, []int{0, 9, 9}, digits)
}

func TestExecute(t *testing.T) {
	instructions := parseInput("inp x\nmul x -1")
	state := InitAluStateFromDigits([]int{10})
	state.Execute(instructions)
	assert.Equal(t, []int{0, -10, 0, 0}, state.variables)

	instructions = parseInput("inp z\ninp x\nmul z 3\neql z x")
	state = InitAluStateFromDigits([]int{2, 6})
	state.Execute(instructions)
	assert.Equal(t, []int{0, 6, 0, 1}, state.variables)
}
