package main

type stack[T any] []T

func New[T any]() *stack[T] {
	return &stack[T]{}
}

func (s *stack[T]) Add(v T) {
	*s = append(*s, v)
}

func (s *stack[T]) Len() int { return len(*s) }
