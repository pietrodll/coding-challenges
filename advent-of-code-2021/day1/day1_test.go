package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseValues(t *testing.T) {
	input := "1\n2\n3\n4\n5"

	assert.Equal(t, []int{1, 2, 3, 4, 5}, parseValues(input))
}

func TestDay1Part1(t *testing.T) {
	values := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	count := countIncrease(values)

	assert.Equal(t, 7, count)
}

func TestDay1Part2(t *testing.T) {
	values := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	expectedGrouped := []int{607, 618, 618, 617, 647, 716, 769, 792}

	grouped := sumBy(values, 3)

	assert.Equal(t, expectedGrouped, grouped)

	count := countIncrease(grouped)

	assert.Equal(t, 5, count)

	assert.Equal(t, []int{}, sumBy([]int{1, 2}, 3))
	assert.Equal(t, 0, countIncrease([]int{}))
}
