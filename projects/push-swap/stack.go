package main

type Stack struct {
	data []int
}

func (s *Stack) push(val int) {
	s.data = append(s.data, val)
}
func (s *Stack) pop() (int, bool) {
	if s.isEmpty() {
		return 0, false
	}
	index := len(s.data) - 1
	topItem := s.data[index]
	s.data = s.data[:index]
	return topItem, true
}
func (s *Stack) peek() (int, bool) {
	if s.isEmpty() {
		return 0, false
	}
	return s.data[len(s.data)-1], true
}
func (s *Stack) isEmpty() bool {
	return len(s.data) == 0
}
func (s *Stack) size() int {
	return len(s.data)
}
