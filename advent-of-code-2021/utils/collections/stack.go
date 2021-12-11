package collections

import (
	"errors"
)

type Stack struct {
	top *linkedElement
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack) Add(element interface{}) {
	toAdd := linkedElement{element, nil}

	if s.IsEmpty() {
		s.top = &toAdd
	} else {
		toAdd.next = s.top
		s.top = &toAdd
	}
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		panic(errors.New("stack is empty, cannot pop"))
	}

	toPop := s.top
	s.top = toPop.next
	return toPop.value
}

func NewStack(elements ...interface{}) Stack {
	s := Stack{nil}

	for _, elem := range elements {
		s.Add(elem)
	}

	return s
}

type IntStack struct {
	s Stack
}

func (s *IntStack) IsEmpty() bool {
	return s.s.IsEmpty()
}

func (s *IntStack) Add(element int) {
	s.s.Add(element)
}

func (s *IntStack) Pop() int {
	return s.s.Pop().(int)
}

func NewIntStack(elements ...int) IntStack {
	s := IntStack{NewStack()}

	for _, element := range elements {
		s.Add(element)
	}

	return s
}
