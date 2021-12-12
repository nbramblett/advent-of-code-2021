package day11

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nbramblett/advent-of-code-2021/gohelp"
)

func Solve1() {
	grid := ReadInput()
	tally := 0
	for i := 0; i < 100; i++ {
		flashPoints := grid.bumpPoints()
		for len(flashPoints) > 0 {
			tally++
			flashPoint := flashPoints[0]
			flashPoints = flashPoints[1:]
			flashPoints = append(flashPoints, grid.flashPoint(flashPoint)...)
		}
	}
	log.Println(tally)
}

type Point struct {
	i, j int
}

type Grid [][]int

func (g Grid) AllZero() bool {
	for _, r := range g {
		for _, c := range r {
			if c > 0 {
				return false
			}
		}
	}
	return true
}

func (g Grid) flashPoint(p Point) []Point {
	newFlashPoints := []Point{}
	i, j := p.i, p.j
	for i2 := i - 1; i2 <= i+1; i2++ {
		for j2 := j - 1; j2 <= j+1; j2++ {
			if (i2 == i && j2 == j) || i2 < 0 || i2 >= len(g) || j2 < 0 || j2 >= len(g[i2]) {
				continue
			}
			if g[i2][j2] > 0 {
				newFlashPoints = append(newFlashPoints, g.bumpPoint(Point{i2, j2})...)
			}
		}
	}
	return newFlashPoints
}

func (g Grid) bumpPoints() []Point {
	flashPoints := []Point{}
	for i := range g {
		for j := range g[i] {
			flashPoints = append(flashPoints, g.bumpPoint(Point{i, j})...)
		}
	}
	return flashPoints
}

func (g Grid) bumpPoint(p Point) []Point {
	flashPoints := []Point{}
	g[p.i][p.j] += 1
	if g[p.i][p.j] > 9 {
		g[p.i][p.j] = 0
		flashPoints = append(flashPoints, p)
	}
	return flashPoints
}

func (g Grid) String() string {
	s := ""
	for i := range g {
		for j := range g[i] {
			s = fmt.Sprintf("%s%d", s, g[i][j])
		}
		s = fmt.Sprintf("%s\n", s)
	}
	return s
}

func ReadInput() Grid {
	file, err := os.Open("day11/input.txt")
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
}
