package grid

import (
	"errors"

	"github.com/pietrodll/aoc2021/utils/base"
)

type GridPoint struct {
	I int
	J int
}

// Implement Codable interface
func (p GridPoint) Encode(c base.Coder) int {
	switch c := c.(type) {
	case *Grid:
		return c.Encode(p)
	default:
		panic(errors.New("GridPoint can only be decoded by Grid"))
	}
}
