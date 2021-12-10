package day10

import (
	"bufio"
	"log"
	"os"
)

var openingRunes = []rune{'(', '{', '[', '<'}

var closingRunes = []rune{')', '}', ']', '>'}

func Solve1() {
	sum := 0
	for _, s := range ReadInput() {
		r := IllegalChar(s)
		switch r {
		case ')':
			sum += 3
		case ']':
			sum += 57
		case '}':
			sum += 1197
		case '>':
			sum += 25137
		}
	}
	log.Println(sum)

}

func IllegalChar(s string) rune {
	openings := []rune{}
	for _, r := range s {
		if isOpening(r) {
			openings = append(openings, r)
		} else if ok, i := isClosing(r); ok {
			matchingOpening := openingRunes[i]
			if openings[len(openings)-1] != matchingOpening {
				return r
			}
			openings = openings[:len(openings)-1]
		}
	}
	return 'n'
}

func isClosing(r rune) (bool, int) {
	for i, r2 := range closingRunes {
		if r == r2 {
			return true, i
		}
	}
	return false, 0
}

func isOpening(r rune) bool {
	for _, r2 := range openingRunes {
		if r == r2 {
			return true
		}
	}
	return false
}

func ReadInput() []string {
	file, err := os.Open("day10/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// first line of input is the numbers to read
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
