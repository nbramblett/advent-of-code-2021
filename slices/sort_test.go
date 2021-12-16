package slices

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSorted(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5, -1, -2, -3, -4}
	intSlice2 := []int{1, 2, 3, 4, 5}
	require.False(t, Sorted(intSlice))
	require.True(t, Sorted(intSlice2))
}

func TestSort(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5, -1, -2, -3, -4}
	Sort(intSlice)
	require.True(t, Sorted(intSlice))
}

type Sortable struct {
	x int
}

func Less(a, b Sortable) bool {
	return a.x < b.x
}

func TestSortF(t *testing.T) {
	intSlice := []Sortable{{-1}, {-5}, {5}, {0}}
	SortF(intSlice, Less)
	require.Equal(t, []Sortable{{-5}, {-1}, {0}, {5}}, intSlice)
}

func TestReverse(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}
	Reverse(intSlice)
	require.Equal(t, []int{5, 4, 3, 2, 1}, intSlice)
}
