package main

type Stack[T any] []any

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Add(v T) {
	*s = append(*s, v)
}
