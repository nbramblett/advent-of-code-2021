package day8

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Solve1() {
	lines := ReadInput()
	tally := 0
	for _, line := range lines {
		end := line[1]
		digits := strings.Split(end, " ")
		for _, digit := range digits {
			if len(digit) == 2 || len(digit) == 3 || len(digit) == 4 || len(digit) == 7 {
				tally++
			}
		}
	}
	log.Println(tally)
}

func ReadInput() [][]string {
	file, err := os.Open("day8/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// first line of input is the numbers to read
	lines := [][]string{}
	for scanner.Scan() {
		lines = append(lines, strings.Split(scanner.Text(), " | "))
	}
	return lines
	panic("no line!")
}
