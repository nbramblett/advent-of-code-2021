package sets

// Set is a generic Set implementation
type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(t ...T) {
	for _, k := range t {
		s[k] = struct{}{}
	}
}

func (s Set[T]) Contains(t T) bool {
	_, ok := s[t]
	return ok
}

func (s Set[T]) Del(t T) {
	delete(s, t)
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) ToSlice() []T {
	sl := make([]T, 0, len(s))
	for k := range s {
		sl = append(sl, k)
	}
	return sl
}

func New[T comparable](ts ...T) Set[T] {
	s := make(Set[T])
	s.Add(ts...)
	return s
}

func Union[T comparable](s1, s2 Set[T]) Set[T] {
	s3 := make(Set[T], len(s1)+len(s2))
	for k, v := range s1 {
		s3[k] = v
	}
	for k, v := range s2 {
		s3[k] = v
	}
	return s3
}

func Intersection[T comparable](s1, s2 Set[T]) Set[T] {
	s3 := make(Set[T])
	for k, v := range s1 {
		if !s2.Contains(k) {
			continue
		}
		s3[k] = v
	}
	return s3
}

func Minus[T comparable](s1, s2 Set[T]) Set[T] {
	s3 := make(Set[T])
	for k, v := range s1 {
		if s2.Contains(k) {
			continue
		}
		s3[k] = v
	}
	return s3
}

// Disjunction returns the disjunctive union of sets.
func Disjunction[T comparable](s1, s2 Set[T]) Set[T] {
	return Minus(Union(s1, s2), Intersection(s1, s2))
}
