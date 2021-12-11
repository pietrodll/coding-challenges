package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntSet(t *testing.T) {
	s := NewIntSet()

	assert.True(t, s.IsEmpty())
	assert.Equal(t, 0, s.Len())
	assert.False(t, s.Contains(0))

	s.Add(0)
	assert.False(t, s.IsEmpty())
	assert.Equal(t, 1, s.Len())
	assert.True(t, s.Contains(0))

	s.Add(1)
	assert.Equal(t, 2, s.Len())

	assert.False(t, s.Contains(2))
	assert.PanicsWithError(t, "element not in set, cannot remove", func() {
		s.Remove(2)
	})

	assert.Contains(t, []int{0, 1}, s.Pop())
	assert.Contains(t, []int{0, 1}, s.Pop())

	assert.Equal(t, 0, s.Len())

	assert.PanicsWithError(t, "set is empty, cannot pop", func() {
		s.Pop()
	})
}
