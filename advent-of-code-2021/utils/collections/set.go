package collections

import "errors"

type IntSet struct {
	container map[int]bool
}

func (s *IntSet) Len() int {
	return len(s.container)
}

func (s *IntSet) IsEmpty() bool {
	return len(s.container) == 0
}

func (s *IntSet) Add(element int) {
	s.container[element] = true
}

func (s *IntSet) Contains(element int) bool {
	_, present := s.container[element]
	return present
}

func (s *IntSet) Pop() int {
	if s.IsEmpty() {
		panic(errors.New("set is empty, cannot pop"))
	}

	var result int

	for element := range s.container {
		result = element
		break
	}

	delete(s.container, result)
	return result
}

func (s *IntSet) Remove(element int) {
	_, present := s.container[element]

	if !present {
		panic(errors.New("element not in set, cannot remove"))
	}

	delete(s.container, element)
}

func NewIntSet(elements ...int) IntSet {
	container := make(map[int]bool)

	for _, element := range elements {
		container[element] = true
	}

	return IntSet{container}
}
