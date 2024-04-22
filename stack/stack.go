package stack

import "errors"

type Stack[T any] struct {
	elements []T
	top      int
}

func (s *Stack[T]) Pop() (T, error) {
	if s.top == -1 {
		return *new(T), errors.New("stack is empty, there's nothing left to pop")
	}
	top := s.elements[s.top]
	s.top--
	return top, nil
}

func (s *Stack[T]) Push(element T) error {
	if s.top+1 >= len(s.elements) {
		return errors.New("stack is full, free space before pushing in new data")
	}
	s.top++
	s.elements[s.top] = element
	return nil
}

func (s *Stack[T]) Peek() (T, error) {
	if s.top < 0 {
		return *new(T), errors.New("stack is empty, there's nothing left to pop")
	}
	return s.elements[s.top], nil
}

func NewStack[T any](size int) *Stack[T] {
	return &Stack[T]{
		top:      -1,
		elements: make([]T, size),
	}
}
