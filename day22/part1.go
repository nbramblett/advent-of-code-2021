package day22

import (
	"log"
	"strings"

	"github.com/nbramblett/advent-of-code-2021/util"
)

func Solve1() {
	lines := util.ReadLines("day22/input.txt")
	onOff, insts := SplitLines(lines)
	points := Eval(onOff, insts)
	tally := 0
	for x := -50; x <= 50; x++ {
		for y := -50; y <= 50; y++ {
			for z := -50; z <= 50; z++ {
				if points[Point{x, y, z}] {
					tally++
				}
			}
		}
	}
	log.Println(tally)
}

func Eval(onOffs []string, insts []string) map[Point]bool {
	points := map[Point]bool{}
	for i := range onOffs {
		min, max := ParseRanges(insts[i])
		for x := min.x; x <= max.x; x++ {
			for y := min.y; y <= max.y; y++ {
				for z := min.z; z <= max.z; z++ {
					if onOffs[i] == "on" {
						points[Point{x, y, z}] = true
					} else {
						points[Point{x, y, z}] = false
					}
				}
			}
		}
	}
	return points
}

type Point struct {
	x, y, z int
}

func SplitLines(lines []string) ([]string, []string) {
	onOffs := []string{}
	coords := []string{}
	for _, l := range lines {
		ls := strings.Split(l, " ")
		onOffs = append(onOffs, ls[0])
		coords = append(coords, ls[1])
	}
	return onOffs, coords
}

func ParseRanges(l string) (Point, Point) {
	coords := strings.Split(l, ",")
	x := util.StringsToInts(strings.Split(coords[0][2:], ".."))
	y := util.StringsToInts(strings.Split(coords[1][2:], ".."))
	z := util.StringsToInts(strings.Split(coords[2][2:], ".."))
	_, x[0] = util.MinMax(x[0], -50)
	x[1], _ = util.MinMax(x[1], 50)
	_, y[0] = util.MinMax(y[0], -50)
	y[1], _ = util.MinMax(y[1], 50)
	_, z[0] = util.MinMax(z[0], -50)
	z[1], _ = util.MinMax(z[1], 50)

	return Point{x[0], y[0], z[0]}, Point{x[1], y[1], z[1]}
}
