package main

type Queue[T any] []T

func (q *Queue[T]) Len() int { return len(*q) }
