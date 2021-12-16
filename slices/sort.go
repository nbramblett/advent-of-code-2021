package slices

import "constraints"

func Sorted[T constraints.Ordered](s []T) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			return false
		}
	}
	return true
}

// Extremely naive bubblesort for now
func Sort[T constraints.Ordered](s []T) {
	for i := range s {
		for j := range s {
			if s[i] < s[j] {
				st := s[i]
				s[i] = s[j]
				s[j] = st
			}
		}
	}
}

// Extremely naive bubblesort for now
func SortF[T any](s []T, less func(T, T) bool) {
	for i := range s {
		for j := range s {
			if less(s[i], s[j]) {
				st := s[i]
				s[i] = s[j]
				s[j] = st
			}
		}
	}
}

func Reverse[T any](s []T) {
	for l, r := 0, len(s)-1; l < r; {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
