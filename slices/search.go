package slices

import "constraints"

// O(n) time
func Contains[T comparable](s []T, k T) bool {
	for _, v := range s {
		if k == v {
			return true
		}
	}
	return false
}

// O(log n) time, only works on ordered
func BinarySearch[T constraints.Ordered](s []T, k T) int {
	a, b := 0, len(s)
	for a < b {
		m := (a + b) / 2
		if s[m] == k {
			return m
		} else if s[m] < k {
			a = m + 1
		} else {
			b = m
		}
	}
	if s[a] != k {
		return -1
	}
	return a
}
