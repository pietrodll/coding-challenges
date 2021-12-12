package collections

import (
	"math"
	"testing"

	"github.com/pietrodll/aoc2021/utils/base"
	"github.com/stretchr/testify/assert"
)

func TestIntSet(t *testing.T) {
	s := NewIntSet()

	assert.True(t, s.IsEmpty())
	assert.Equal(t, 0, s.Len())
	assert.False(t, s.Contains(0))

	s = NewIntSet(0)
	assert.False(t, s.IsEmpty())
	assert.Equal(t, 1, s.Len())
	assert.True(t, s.Contains(0))

	s.Add(1)
	assert.Equal(t, 2, s.Len())

	s.Add(2)
	assert.True(t, s.Contains(2))

	s.Remove(2)
	assert.False(t, s.Contains(2))
	assert.PanicsWithError(t, "element not in set, cannot remove", func() {
		s.Remove(2)
	})

	assert.ElementsMatch(t, []int{0, 1}, s.ToArray())
	assert.Contains(t, []int{0, 1}, s.Pop())
	assert.Contains(t, []int{0, 1}, s.Pop())

	assert.Equal(t, 0, s.Len())

	assert.PanicsWithError(t, "set is empty, cannot pop", func() {
		s.Pop()
	})
}

// Struct that implements the Codable interface
type Point struct {
	X int
	Y int
}

func (p Point) Encode(coder base.Coder) int {
	return coder.Encode(p)
}

type PointEncoder struct{}

func (p *PointEncoder) Encode(point base.Codable) int {
	switch point := point.(type) {
	case Point:
		return int(math.Pow(2, float64(point.X))) * (2*point.Y + 1)
	default:
		return 0
	}
}

func (p *PointEncoder) Decode(encoded int) base.Codable {
	x := 0
	for encoded%2 == 0 {
		x++
		encoded /= 2
	}

	y := encoded / 2
	return Point{x, y}
}

func TestCodableSet(t *testing.T) {
	coder := PointEncoder{}

	s := NewCodableSet(&coder, Point{1, 1})
	assert.Equal(t, 1, s.Len())
	assert.False(t, s.IsEmpty())
	assert.True(t, s.Contains(Point{1, 1}))

	s.Add(Point{5, 6})
	assert.ElementsMatch(t, []Point{{1, 1}, {5, 6}}, s.ToArray())
	s.Remove(Point{5, 6})
	assert.Equal(t, Point{1, 1}, s.Pop())
}

func TestStringSet(t *testing.T) {
	s := NewStringSet()

	assert.True(t, s.IsEmpty())
	assert.Equal(t, 0, s.Len())
	assert.False(t, s.Contains("a"))

	s = NewStringSet("a")
	assert.False(t, s.IsEmpty())
	assert.Equal(t, 1, s.Len())
	assert.True(t, s.Contains("a"))

	s.Add("b")
	assert.Equal(t, 2, s.Len())

	s.Add("c")
	assert.True(t, s.Contains("c"))

	s.Remove("c")
	assert.False(t, s.Contains("c"))
	assert.PanicsWithError(t, "element not in set, cannot remove", func() {
		s.Remove("c")
	})

	assert.ElementsMatch(t, []string{"a", "b"}, s.ToArray())

	assert.Contains(t, []string{"a", "b"}, s.Pop())
	assert.Contains(t, []string{"a", "b"}, s.Pop())

	assert.Equal(t, 0, s.Len())

	assert.PanicsWithError(t, "set is empty, cannot pop", func() {
		s.Pop()
	})
}
