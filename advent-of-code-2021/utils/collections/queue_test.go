package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntQueue(t *testing.T) {
	q := NewIntQueue()

	assert.Nil(t, q.head)
	assert.Nil(t, q.tail)

	q.Enqueue(1)

	assert.Equal(t, q.head.value, 1)
	assert.Nil(t, q.head.next)
	assert.Equal(t, q.head, q.tail)

	q.Enqueue(2)
	q.Enqueue(3)

	assert.Equal(t, q.head.value, 1)
	assert.Equal(t, q.tail.value, 3)
	assert.Equal(t, q.head.next.next, q.tail)
	assert.Nil(t, q.tail.next)

	val := q.Dequeue()

	assert.Equal(t, 1, val)
	assert.Equal(t, 2, q.head.value)
	assert.Equal(t, q.head.next, q.tail)
	assert.Equal(t, q.tail.value, 3)

	q.Dequeue()
	q.Dequeue()

	assert.True(t, q.IsEmpty())

	assert.Panics(t, func() {
		q.Dequeue()
	})
}
