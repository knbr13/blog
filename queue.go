package main

import (
	"fmt"
	"strings"
)

type queue[T any] []T

func NewQueue[T any]() *queue[T] {
	return &queue[T]{}
}

func (q *queue[T]) String() string {
	var sb strings.Builder

	sb.WriteString(strings.Repeat("—", q.MaxLen()+4))
	sb.WriteString("\n")

	for i := len(*q) - 1; i >= 0; i-- {
		sb.WriteString("|")
		sb.WriteString(fmt.Sprintf(" %v ", (*q)[i]))
		sb.WriteString("|\n")
	}

	sb.WriteString(strings.Repeat("—", q.MaxLen()+4))
	sb.WriteString("\n")

	return sb.String()
}

func (q *queue[T]) Len() int { return len(*q) }

func (q *queue[T]) Enqueue(t T) { *q = append(*q, t) }

func (q *queue[T]) Dequeue() *T {
	if q.Len() == 0 {
		return nil
	}
	deleted := (*q)[0]
	(*q) = (*q)[1:]
	return &deleted
}

func (q *queue[T]) Peek() *T {
	if q.Len() == 0 {
		return nil
	}
	return &(*q)[0]
}

func (q *queue[T]) Empty() bool {
	return q.Len() == 0
}

func (q *queue[T]) Size() int {
	return q.Len()
}

func (q *queue[T]) Clear() {
	*q = (*q)[:0]
}

func (q *queue[T]) MaxLen() int {
	max := ""
	for _, v := range *q {
		str := fmt.Sprintf("%v", v)
		if len(str) > len(max) {
			max = str
		}
	}
	return len(max)
}
