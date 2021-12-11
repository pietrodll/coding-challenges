package collections

import (
	"errors"
)

type linkedElement struct {
	value interface{}
	next  *linkedElement
}

type Queue struct {
	head *linkedElement
	tail *linkedElement
}

func (q *Queue) IsEmpty() bool {
	return q.head == nil
}

func (q *Queue) Enqueue(element interface{}) {
	newElement := linkedElement{element, nil}

	if q.IsEmpty() {
		q.head = &newElement
		q.tail = &newElement
	} else {
		q.tail.next = &newElement
		q.tail = &newElement
	}
}

func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		panic(errors.New("queue is empty, cannot dequeue"))
	}

	head := q.head
	q.head = head.next

	return head.value
}

func NewQueue(elements ...interface{}) Queue {
	q := Queue{nil, nil}

	for _, elem := range elements {
		q.Enqueue(elem)
	}

	return q
}

type IntQueue struct {
	q Queue
}

func (q *IntQueue) IsEmpty() bool {
	return q.q.IsEmpty()
}

func (q *IntQueue) Enqueue(element int) {
	q.q.Enqueue(element)
}

func (q *IntQueue) Dequeue() int {
	return q.q.Dequeue().(int)
}

func NewIntQueue(elements ...int) IntQueue {
	q := IntQueue{NewQueue()}

	for _, element := range elements {
		q.Enqueue(element)
	}

	return q
}
