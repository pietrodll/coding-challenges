package collections

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestElement struct {
	key      string
	priority *map[string]int
}

func (e TestElement) GetPriority() int {
	return (*e.priority)[e.key]
}

func (e TestElement) Key() interface{} {
	return e.key
}

func TestPriorityQueue(t *testing.T) {
	priority := make(map[string]int)

	for i := 0; i < 10; i++ {
		priority[fmt.Sprint(i)] = i
	}

	q := NewPriorityQueue()
	assert.True(t, q.IsEmpty())

	q = NewPriorityQueue(TestElement{"3", &priority}, TestElement{"5", &priority}, TestElement{"1", &priority})
	assert.False(t, q.IsEmpty())
	assert.Equal(t, 3, q.Len())
	assert.Equal(t, TestElement{"1", &priority}, q.GetMin())
	assert.Equal(t, TestElement{"1", &priority}, q.Pop())
	assert.Equal(t, 2, q.Len())

	q.Add(TestElement{"4", &priority})
	q.Add(TestElement{"2", &priority})
	q.Add(TestElement{"6", &priority})
	assert.Equal(t, TestElement{"2", &priority}, q.Pop())
	assert.Equal(t, TestElement{"3", &priority}, q.Pop())

	assert.Equal(t, TestElement{"4", &priority}, q.GetMin())

	priority["4"] = 15
	q.IncreasedPriority(TestElement{"4", &priority})
	assert.Equal(t, TestElement{"5", &priority}, q.GetMin())

	priority["4"] = 0
	q.DecreasedPriority(TestElement{"4", &priority})
	assert.Equal(t, TestElement{"4", &priority}, q.GetMin())
}
