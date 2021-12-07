package day7

import (
	"log"
	"math"
)

func Solve2() {
	vals := ReadInput()
	distanceMetric = func(x, y int) int {
		n := int(math.Abs(float64(x - y)))
		return n * (n + 1) / 2
	}
	od := optimizeDistance(vals)
	log.Println(od)
}
