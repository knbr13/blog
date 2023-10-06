package main

type Queue[T any] []T

func (q *Queue[T]) Len() int { return len(*q) }

func (q *Queue[T]) Add(t T) { *q = append(*q, t) }
