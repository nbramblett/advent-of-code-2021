package gohelp

import "strconv"

func StringsToInts(strs []string) []int {
	ints := []int{}
	for _, i := range strs {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		ints = append(ints, j)
	}
	return ints
}

func MinMax(vars ...int) (int, int) {
	min := vars[0]
	max := min

	for _, i := range vars {
		if min > i {
			min = i
		}
		if max < i {
			max = i
		}
	}

	return min, max
}
