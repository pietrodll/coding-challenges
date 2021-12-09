package collections

import (
	"errors"
)

type intLinkedElement struct {
	value int
	next  *intLinkedElement
}

type Queue interface {
	IsEmpty() bool
	Enqueue(element interface{})
	Dequeue() interface{}
}

type IntQueue struct {
	head *intLinkedElement
	tail *intLinkedElement
}

func (q *IntQueue) IsEmpty() bool {
	return q.head == nil
}

func (q *IntQueue) Enqueue(element int) {
	new_element := intLinkedElement{element, nil}

	if q.IsEmpty() {
		q.head = &new_element
		q.tail = &new_element
	} else {
		q.tail.next = &new_element
		q.tail = &new_element
	}
}

func (q *IntQueue) Dequeue() int {
	if q.IsEmpty() {
		panic(errors.New("queue is empty, cannot dequeue"))
	}

	head := q.head
	q.head = head.next

	return head.value
}

func NewIntQueue(elements ...int) IntQueue {
	q := IntQueue{nil, nil}

	for _, elem := range elements {
		q.Enqueue(elem)
	}

	return q
}
