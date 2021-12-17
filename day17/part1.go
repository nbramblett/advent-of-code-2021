package day17

import (
	"bufio"
	"log"
	"os"

	"github.com/nbramblett/advent-of-code-2021/sets"
)

// Solves 1 and 2
func Solve1() {
	targetXRange := Range{195, 238}
	targetYRange := Range{-93, -67}
	// test values
	// targetXRange := Range{20, 30}
	// targetYRange := Range{-10, -5}

	ranges := map[int]Range{}
	for velX := 0; velX <= 238; velX++ {
		r := Range{}
		addRange := false
		for n := 0; n <= targetXRange.high; n++ {
			val := XAfterNSteps(velX, n)
			if targetXRange.InRange(val) {
				addRange = true
				if r.low == 0 {
					r.low = n
				} else {
					r.high = n
				}
			} else if val > targetXRange.high {
				break
			}
		}
		if r.high == 0 {
			r.high = r.low
		}
		if addRange {
			ranges[velX] = r
		}
	}
	combos := sets.Set[Range]{}
	yVals := sets.Set[int]{}
	for y := -500; y < 1000; y++ {
	MIDDLE:
		for velX, r := range ranges {
			for n := r.low; n <= r.high; n++ {
				if targetYRange.InRange(YAfterNSteps(y, n)) {
					combos.Add(Range{velX, y})
					yVals.Add(y)
					continue MIDDLE
				}
			}
		}
	}

	bestAlt := -5
	for yVal := range yVals {
		for n := 0; n < 500; n++ {
			y := YAfterNSteps(yVal, n)
			if bestAlt <= y {
				bestAlt = y
			}
		}

	}
	log.Println(bestAlt)
	log.Println(len(combos))
}

type Range struct {
	low, high int
}

func (r Range) InRange(n int) bool {
	return n >= r.low && n <= r.high
}

func XAfterNSteps(svx, n int) int {
	if n > svx {
		n = svx
	}
	return svx*(svx+1)/2 - (svx-n)*(svx-n+1)/2
}

func YAfterNSteps(svy, n int) int {
	return svy*(svy+1)/2 - (svy-n)*(svy-n+1)/2
}

func ReadLines() []string {
	file, err := os.Open("day17/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// first line of input is the numbers to read
	bits := []string{}
	for scanner.Scan() {
		bits = append(bits, scanner.Text())
	}
	return bits
}
