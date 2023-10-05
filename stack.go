package main

import (
	"fmt"
	"strings"
)

type stack[T any] []T

func (s *stack[T]) String() string {
	var sb strings.Builder

	sb.WriteString(strings.Repeat("—", s.MaxLen()+4))
	sb.WriteString("\n")

	for i := len(*s) - 1; i >= 0; i-- {
		sb.WriteString("|")
		sb.WriteString(fmt.Sprintf(" %v ", (*s)[i]))
		sb.WriteString("|\n")
	}

	sb.WriteString(strings.Repeat("—", s.MaxLen()+4))
	sb.WriteString("\n")

	return sb.String()
}

func New[T any]() *stack[T] {
	return &stack[T]{}
}

func (s *stack[T]) Push(v T) {
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
	return &(*s)[s.Len()-1]
}

func (s *stack[T]) Clear() {
	*s = (*s)[:0]
}

func (s *stack[T]) Empty() bool {
	return s.Len() == 0
}

func (s *stack[T]) MaxLen() int {
	max := ""
	for _, v := range *s {
		str := fmt.Sprintf("%v", v)
		if len(str) > len(max) {
			max = str
		}
	}
	return len(max)
}
