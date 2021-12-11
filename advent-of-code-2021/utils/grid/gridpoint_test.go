package grid

import (
	"testing"

	"github.com/pietrodll/aoc2021/utils/base"
	"github.com/stretchr/testify/assert"
)

type DummyCoder struct{}

func (d *DummyCoder) Encode(c base.Codable) int {
	return 0
}

func (d *DummyCoder) Decode(encoded int) base.Codable {
	return nil
}

func TestGridPointEncoding(t *testing.T) {
	p := GridPoint{0, 0}
	grid := NewGrid([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	assert.Equal(t, 0, p.Encode(&grid))

	assert.PanicsWithError(t, "GridPoint can only be decoded by Grid", func() {
		p.Encode(&DummyCoder{})
	})
}
