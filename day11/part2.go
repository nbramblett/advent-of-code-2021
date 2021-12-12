package day11

import "log"

func Solve2() {
	grid := ReadInput()
	tally := 0
	for !grid.AllZero() {
		tally++
		flashPoints := grid.bumpPoints()
		for len(flashPoints) > 0 {
			flashPoint := flashPoints[0]
			flashPoints = flashPoints[1:]
			flashPoints = append(flashPoints, grid.flashPoint(flashPoint)...)
		}
	}
	log.Println(tally)
}
