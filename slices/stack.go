package slices

type Stack[T any] []T

func (s *Stack[T]) Push(t ...T) {
	*s = append(*s, t...)
}

func (s *Stack[T]) Pop() T {
	t := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return t
}
