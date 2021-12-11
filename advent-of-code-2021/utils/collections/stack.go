package collections

import "errors"

type IntStack struct {
	top *intLinkedElement
}

func (s *IntStack) IsEmpty() bool {
	return s.top == nil
}

func (s *IntStack) Add(element int) {
	toAdd := intLinkedElement{element, nil}

	if s.IsEmpty() {
		s.top = &toAdd
	} else {
		toAdd.next = s.top
		s.top = &toAdd
	}
}

func (s *IntStack) Pop() int {
	if s.IsEmpty() {
		panic(errors.New("stack is empty, cannot pop"))
	}

	toPop := s.top
	s.top = toPop.next
	return toPop.value
}

func NewIntStack(elements ...int) IntStack {
	s := IntStack{nil}

	for _, elem := range elements {
		s.Add(elem)
	}

	return s
}
