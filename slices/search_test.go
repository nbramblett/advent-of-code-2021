package slices

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContains(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5, -1, -2, -3, -4}
	require.True(t, Contains(intSlice, 5))
}

func TestBinarySearch(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5, 6, 7}
	require.Equal(t, 5, BinarySearch(intSlice, 6))
	require.Equal(t, -1, BinarySearch(intSlice, 12))
}
