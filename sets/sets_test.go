package sets

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSet(t *testing.T) {
	set := Set[int]{}

	set.Add(1)
	require.Equal(t, 1, set.Len())
	require.True(t, set.Contains(1))
	set.Del(1)
	require.Equal(t, 0, set.Len())
	set.Add(1, 2, 3)
	require.Equal(t, 3, set.Len())

	require.ElementsMatch(t, []int{1, 2, 3}, set.ToSlice())
}

func TestFunctions(t *testing.T) {
	set1 := New(1, 2, 3)
	set2 := New(2, 3, 4)

	s := Union(set1, set2)
	require.Equal(t, New(1, 2, 3, 4), s)

	s = Intersection(set1, set2)
	require.Equal(t, New(2, 3), s)

	s = Minus(set1, set2)
	require.Equal(t, New(1), s)

	s = Disjunction(set1, set2)
	require.Equal(t, New(1, 4), s)
}
