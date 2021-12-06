package day5

import (
	"fmt"
	"math"
)

func Solve2() {
	l := Line{0, 5, 0, -5}
	fmt.Println(l.CoverWithDiag())
	lines := ReadInput()
	coverPoints := map[Point]int{}
	for _, line := range lines {
		points := line.CoverWithDiag()
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

func (l Line) CoverWithDiag() []Point {
	points := l.Cover()
	if points != nil {
		return points
	}
	for i := 0; i <= int(math.Abs(float64(l.x2-l.x1))); i++ {
		dx, dy := i, i
		if l.x2 < l.x1 {
			dx = -i
		}
		if l.y2 < l.y1 {
			dy = -i
		}
		points = append(points, Point{l.x1 + dx, l.y1 + dy})
	}
	return points
}
