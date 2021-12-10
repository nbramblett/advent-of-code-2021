package day10

import (
	"log"
	"sort"
)

func Solve2() {
	sums := []int{}
	for _, s := range ReadInput() {
		r := IllegalChar(s)
		if r != 'n' {
			continue
		}
		cs := IncompleteChars(s)
		sum := 0
		for _, r2 := range cs {
			sum *= 5
			switch r2 {
			case '(':
				sum += 1
			case '[':
				sum += 2
			case '{':
				sum += 3
			case '<':
				sum += 4
			}
		}
		sums = append(sums, sum)
	}
	sort.Ints(sums)

	log.Println(sums[len(sums)/2])

}

func IncompleteChars(s string) []rune {
	openings := []rune{}
	for _, r := range s {
		if isOpening(r) {
			openings = append(openings, r)
		} else if ok, _ := isClosing(r); ok {
			openings = openings[:len(openings)-1]
		}
	}
	// Reverse openings!
	for i, j := 0, len(openings)-1; i < j; i, j = i+1, j-1 {
		openings[i], openings[j] = openings[j], openings[i]
	}

	return openings
}
