package main

type Queue[T any] []T

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Len() int { return len(*q) }

func (q *Queue[T]) Enqueue(t T) { *q = append(*q, t) }

func (q *Queue[T]) Dequeue() *T {
	if q.Len() == 0 {
		return nil
	}
	deleted := (*q)[0]
	(*q) = (*q)[1:]
	return &deleted
}
