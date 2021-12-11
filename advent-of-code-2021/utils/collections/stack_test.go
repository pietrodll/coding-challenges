package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := NewIntStack()

	assert.Nil(t, s.top)

	s.Add(1)
	s.Add(2)

	assert.Equal(t, 2, s.top.value)
	assert.Equal(t, 1, s.top.next.value)

	val := s.Pop()

	assert.Equal(t, 2, val)
	assert.Equal(t, 1, s.top.value)
	assert.Nil(t, s.top.next)
}
