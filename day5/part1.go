package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nbramblett/advent-of-code-2021/gohelp"
)

func Solve1() {
	lines := ReadInput()
	coverPoints := map[Point]int{}
	for _, line := range lines {
		points := line.Cover()
		for _, point := range points {
			coverPoints[point] += 1
		}
	}
	count := 0
	for _, total := range coverPoints {
		if total >= 2 {
			count++
		}
	}
	fmt.Println(len(lines))
	fmt.Println(len(coverPoints))
	fmt.Println(count)
}

type Point struct {
	x, y int
}

type Line struct {
	x1, x2, y1, y2 int
}

func (l Line) Cover() []Point {
	// Handle Vert and Hor cases only (for now)
	if !l.IsOrthogonal() {
		return nil
	}
	points := []Point{}
	if l.x1 == l.x2 {
		start, end := gohelp.MinMax(l.y1, l.y2)
		for i := start; i <= end; i++ {
			points = append(points, Point{l.x1, i})
		}
	} else {
		start, end := gohelp.MinMax(l.x1, l.x2)
		for i := start; i <= end; i++ {
			points = append(points, Point{i, l.y1})
		}
	}
	return points
}

func (l Line) IsOrthogonal() bool {
	return l.x1 == l.x2 || l.y1 == l.y2
}

func ReadInput() []Line {
	file, err := os.Open("day5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := []Line{}
	scanner := bufio.NewScanner(file)
	// first line of input is the numbers to read
	for scanner.Scan() {
		line := scanner.Text()
		coords := strings.Split(line, " -> ")
		if len(coords) != 2 {
			continue
		}
		start := gohelp.StringsToInts(strings.Split(coords[0], ","))
		end := gohelp.StringsToInts(strings.Split(coords[1], ","))
		lines = append(lines, Line{x1: start[0], y1: start[1], x2: end[0], y2: end[1]})
	}
	return lines
}
