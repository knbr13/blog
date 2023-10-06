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

func (q *Queue[T]) Peek() *T {
	if q.Len() == 0 {
		return nil
	}
	return &(*q)[0]
}

func (q *Queue[T]) Empty() bool {
	return q.Len() == 0
}

func (q *Queue[T]) Size() int {
	return q.Len()
}

func (q *Queue[T]) Clear() {
	*q = (*q)[:0]
}
