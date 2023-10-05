package main

type stack[T any] []T

func New[T any]() *stack[T] {
	return &stack[T]{}
}

func (s *stack[T]) Add(v T) {
	*s = append(*s, v)
}

func (s *stack[T]) Len() int { return len(*s) }

func (s *stack[T]) Pop() *T {
	if s.Len() == 0 {
		return nil
	}
	deleted := (*s)[s.Len()-1]
	*s = (*s)[:s.Len()-1]
	return &deleted
}

func (s *stack[T]) Peek() *T {
	if s.Len() == 0 {
		return nil
	}
	return &(*s)[s.Len()]
}

func (s *stack[T]) Clear() {
	*s = (*s)[:0]
}
