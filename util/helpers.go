package util

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

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

func ReadLines(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// first line of input is the numbers to read
	nodes := []string{}
	for scanner.Scan() {
		nodes = append(nodes, scanner.Text())
	}
	return nodes
}
