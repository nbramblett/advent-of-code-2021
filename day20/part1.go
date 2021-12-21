package day20

import (
	"log"
	"math"
	"strconv"

	"github.com/nbramblett/advent-of-code-2021/util"
)

var theRule rule

// everything outside defaults to dark, meaning it will evaluate as the rule[0]. After that, everything is either
// ., in which case it stays ., or it's #, in which case we evaluate rule[-1] and it's either back to . or it stays on #
var atInfinity rune

func Solve1() {
	lines := util.ReadLines("day20/input.txt")
	r, i := ReadInput(lines)
	theRule = r
	atInfinity = '.'

	i2 := Step(Step(i))
	log.Println(i2.tally('#'))
}

func Solve2() {
	lines := util.ReadLines("day20/input.txt")
	r, i := ReadInput(lines)
	theRule = r
	atInfinity = '.'

	for n := 0; n < 50; n++ {
		i = Step(i)
	}
	log.Println(i.tally('#'))
}

func Step(i image) image {
	min, max := i.bounds()
	newImg := image{}
	for x := min.x - 2; x <= max.x+2; x++ {
		for y := min.y - 2; y <= max.y+2; y++ {
			r := i.pixelValue(Point{x, y})
			newImg[Point{x, y}] = r
		}
	}
	if atInfinity == '#' {
		atInfinity = rune(theRule[len(theRule)-1])
	} else {
		atInfinity = rune(theRule[0])
	}
	return newImg
}

type rule string

type Point struct{ x, y int }

type image map[Point]rune

func (i image) tally(s rune) (count int) {
	for _, r := range i {
		if r == s {
			count++
		}
	}
	return count
}

func (img image) pixelValue(p Point) rune {
	binaryString := ""
	for j := p.y - 1; j <= p.y+1; j++ {
		for i := p.x - 1; i <= p.x+1; i++ {
			c, ok := img[Point{i, j}]
			if !ok {
				c = atInfinity
			}
			if c == '.' {
				binaryString += "0"
			} else {
				binaryString += "1"
			}
		}
	}
	index, err := strconv.ParseInt(binaryString, 2, 32)
	util.PanicIf(err)
	return rune(theRule[int(index)])
}

func (i image) bounds() (Point, Point) {
	minX, minY := math.MaxInt32, math.MaxInt32
	maxX, maxY := math.MinInt64, math.MinInt64
	for p := range i {
		if p.x > maxX {
			maxX = p.x
		}
		if p.x < minX {
			minX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
		if p.y < minY {
			minY = p.y
		}
	}

	return Point{minX, minY}, Point{maxX, maxY}
}

func (i image) String() string {
	s := "\n"
	min, max := i.bounds()
	for y := min.y; y <= max.y; y++ {
		for j := min.x; j <= max.x; j++ {
			s += string(i[Point{j, y}])
		}
		s += "\n"
	}
	return s
}

func ReadInput(lines []string) (rule, image) {
	r := rule(lines[0])
	i := image{}
	lines = lines[2:]

	for y := range lines {
		for x, r := range lines[y] {
			i[Point{x, y}] = r
		}
	}
	return r, i
}
