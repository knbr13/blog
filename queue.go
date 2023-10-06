package main

import "fmt"

type queue[T any] []T

func Newqueue[T any]() *queue[T] {
	return &queue[T]{}
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
