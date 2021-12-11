package collections

import (
	"errors"

	"github.com/pietrodll/aoc2021/utils/base"
)

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

type CodableSet struct {
	s     IntSet
	coder base.Coder
}

func (s *CodableSet) Len() int {
	return s.s.Len()
}

func (s *CodableSet) IsEmpty() bool {
	return s.s.IsEmpty()
}

func (s *CodableSet) Add(element base.Codable) {
	s.s.Add(s.coder.Encode(element))
}

func (s *CodableSet) Contains(element base.Codable) bool {
	return s.s.Contains(s.coder.Encode(element))
}

func (s *CodableSet) Pop() base.Codable {
	return s.coder.Decode(s.s.Pop())
}

func (s *CodableSet) Remove(element base.Codable) {
	s.s.Remove(s.coder.Encode(element))
}

func NewCodableSet(coder base.Coder, elements ...base.Codable) CodableSet {
	s := NewIntSet()

	for _, element := range elements {
		s.Add(coder.Encode(element))
	}

	return CodableSet{s, coder}
}
