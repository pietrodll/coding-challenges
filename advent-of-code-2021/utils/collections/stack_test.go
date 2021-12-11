package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := NewStack()

	assert.Nil(t, s.top)

	s.Add(1)
	s.Add(2)

	assert.Equal(t, 2, s.top.value)
	assert.Equal(t, 1, s.top.next.value)

	val := s.Pop()

	assert.Equal(t, 2, val)
	assert.Equal(t, 1, s.top.value)
	assert.Nil(t, s.top.next)

	s = NewStack(1, 2)
	assert.Equal(t, 2, s.Pop())
	assert.Equal(t, 1, s.Pop())
	assert.PanicsWithError(t, "stack is empty, cannot pop", func() {
		s.Pop()
	})
}

func TestIntStack(t *testing.T) {
	s := NewIntStack(42)

	assert.Equal(t, s.s.top.value, 42)

	assert.Equal(t, 42, s.Pop())
	assert.True(t, s.IsEmpty())
}
