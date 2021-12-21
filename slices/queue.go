package slices

type Queue[T any] []T

func (q *Queue[T]) Push(t ...T) {
	*q = append(*q, t...)
}

func (q *Queue[T]) Pop() T {
	qv := *q
	t := qv[0]
	*q = qv[1:]
	return t
}
