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

func (s *IntSet) ToArray() []int {
	arr := make([]int, 0)

	for element := range s.container {
		arr = append(arr, element)
	}

	return arr
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

func (s *CodableSet) ToArray() []base.Codable {
	arr := make([]base.Codable, 0)

	for encoded := range s.s.container {
		arr = append(arr, s.coder.Decode(encoded))
	}

	return arr
}

func NewCodableSet(coder base.Coder, elements ...base.Codable) CodableSet {
	s := NewIntSet()

	for _, element := range elements {
		s.Add(coder.Encode(element))
	}

	return CodableSet{s, coder}
}

type StringSet struct {
	container map[string]bool
}

func (s *StringSet) Len() int {
	return len(s.container)
}

func (s *StringSet) IsEmpty() bool {
	return len(s.container) == 0
}

func (s *StringSet) Add(element string) {
	s.container[element] = true
}

func (s *StringSet) Contains(element string) bool {
	_, present := s.container[element]
	return present
}

func (s *StringSet) Pop() string {
	if s.IsEmpty() {
		panic(errors.New("set is empty, cannot pop"))
	}

	var result string

	for element := range s.container {
		result = element
		break
	}

	delete(s.container, result)
	return result
}

func (s *StringSet) Remove(element string) {
	_, present := s.container[element]

	if !present {
		panic(errors.New("element not in set, cannot remove"))
	}

	delete(s.container, element)
}

func (s *StringSet) ToArray() []string {
	arr := make([]string, 0)

	for element := range s.container {
		arr = append(arr, element)
	}

	return arr
}

func NewStringSet(elements ...string) StringSet {
	container := make(map[string]bool)

	for _, element := range elements {
		container[element] = true
	}

	return StringSet{container}
}
