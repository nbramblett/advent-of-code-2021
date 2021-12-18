package main

import (
	"log"
	"time"

	"github.com/nbramblett/advent-of-code-2021/day18"
)

func main() {
	start := time.Now()
	day18.Solve1()
	day18.Solve2()
	elapsed := time.Since(start)
	log.Printf("Day 18 took %s", elapsed)
}
