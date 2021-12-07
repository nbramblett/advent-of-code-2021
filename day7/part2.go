package day7

import (
	"log"
	"math"
	"time"
)

func Solve2() {
	vals := ReadInput()
	now := time.Now()
	distanceMetric = func(x, y int) int {
		n := int(math.Abs(float64(x - y)))
		return n * (n + 1) / 2
	}
	od := optimizeDistance(vals)
	elapsed := time.Since(now)
	log.Println(od)
	log.Printf("Day 6 part 2 actual calculations took %s", elapsed)
}
