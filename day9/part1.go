package day9

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/nbramblett/advent-of-code-2021/gohelp"
)

func Solve1() {
	grid := ReadInput()
	lowPoints := LowPoints(grid)
	sum := 0
	for _, v := range lowPoints {
		sum += v + 1
	}
	log.Println(sum)
}

func LowPoints(grid [][]int) []int {
	lowVals := []int{}
	for i := range grid {
		for k, v := range grid[i] {
			lowest := true
			if k > 0 && grid[i][k-1] <= v {
				lowest = false
			}
			if k < len(grid[i])-1 && grid[i][k+1] <= v {
				lowest = false
			}
			if i > 0 && grid[i-1][k] <= v {
				lowest = false
			}
			if i < len(grid)-1 && grid[i+1][k] <= v {
				lowest = false
			}
			if lowest {
				lowVals = append(lowVals, v)
			}
		}
	}
	return lowVals
}

func ReadInput() [][]int {
	file, err := os.Open("day9/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// first line of input is the numbers to read
	lines := [][]int{}
	for scanner.Scan() {
		lines = append(lines, gohelp.StringsToInts(strings.Split(scanner.Text(), "")))
	}
	return lines
	panic("no line!")
}
